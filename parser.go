package torrentparser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var (
	removeNonDigits *regexp.Regexp
)

func init() {
	var err error
	removeNonDigits, err = regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatalln(err)
	}
}

type ContentType int64

const (
	TV ContentType = iota
	Movie
	Unknown
)

func (c ContentType) String() string {
	switch c {
	case TV:
		return "TV"
	case Movie:
		return "Movie"
	default:
		return "Unknown"
	}
}

type Torrent struct {
	Title       string      `json:"title"`
	ContentType ContentType `json:"contentType"`
	Year        int         `json:"year"`
	Resolution  string      `json:"resolution"`
	Extended    bool        `json:"extended"`
	Unrated     bool        `json:"unrated"`
	Proper      bool        `json:"proper"`
	Repack      bool        `json:"repack"`
	Convert     bool        `json:"convert"`
	Hardcoded   bool        `json:"hardcoded"`
	Retail      bool        `json:"retail"`
	Remastered  bool        `json:"remastered"`
	Region      string      `json:"region"`
	Container   string      `json:"container"`
	Source      string      `json:"source"`
	Codec       string      `json:"codec"`
	Audio       string      `json:"audio"`
	Group       string      `json:"group"`
	Season      int         `json:"season"`
	Episode     int         `json:"episode"`
	Language    string      `json:"language"`
	Hdr         bool        `json:"hdr"`
	ColorDepth  string      `json:"colorDepth"`
	Date        string      `json:"date"`
}

type Parser struct {
	Name            string
	MatchedIndicies map[string]Index
	LowestIndex     int
	LowestWasZero   bool
}

type Index struct {
	Name  string
	Start int
	End   int
}

func ParseName(name string) *Parser {
	return &Parser{Name: name, MatchedIndicies: map[string]Index{}, LowestIndex: len(name)}
}

func (p *Parser) Parse() (Torrent, error) {
	torrent := Torrent{
		Season: -1,
	}

	torrent.Date = p.GetDate()
	torrent.Year = p.GetYear()
	torrent.Container = p.GetContainer()
	torrent.Group = p.GetGroup()
	torrent.Hardcoded = p.GetHardcoded()
	torrent.Remastered = p.GetRemastered()
	torrent.Region = p.GetRegion()
	torrent.Source = p.GetSource()
	torrent.Codec = p.GetCodec()
	torrent.Audio = p.GetAudio()
	torrent.Season = p.GetSeason()
	torrent.Episode = p.GetEpisode()
	torrent.Unrated = p.GetUnrated()
	torrent.Hdr = p.GetHdr()
	torrent.ColorDepth = p.GetColorDepth()
	torrent.Resolution = p.GetResolution()
	torrent.Language = p.GetLanguage()

	// Workaround for checking if episode is part of title
	yearIndex, yearOk := p.MatchedIndicies["year"]
	episodeIndex, epOk := p.MatchedIndicies["episode"]
	if yearOk && epOk && yearIndex.Start > episodeIndex.Start && torrent.Season == -1 {
		torrent.Episode = 0
		p.LowestIndex = yearIndex.Start
		delete(p.MatchedIndicies, "episode")
	}

	// LAST
	torrent.Title = p.GetTitle()

	if torrent.Episode > 0 || torrent.Date != "" || torrent.Season > -1 {
		torrent.ContentType = TV
	} else if torrent.Season == -1 && torrent.Episode == 0 {
		torrent.ContentType = Movie
	}

	return torrent, nil
}

func (p *Parser) MatchedRange(start, end int) bool {
	for _, aRange := range p.MatchedIndicies {
		if start <= aRange.Start && end >= aRange.End {
			return true
		}
		if start >= aRange.Start && start < aRange.End {
			return true
		}
		if end > aRange.Start && end <= aRange.End {
			return true
		}
	}
	return false
}

func (p *Parser) AddMatchedIndex(attr string, loc []int) {
	p.MatchedIndicies[attr] = Index{Name: attr, Start: loc[0], End: loc[1]}
}

func (p *Parser) FindBoolean(attr string, rx *regexp.Regexp) bool {
	loc := rx.FindStringIndex(p.Name)

	if len(loc) == 0 {
		return false
	}
	if p.MatchedRange(loc[0], loc[1]) {
		return false
	}
	p.setLowestIndex(loc[0])
	return true
}

type FindStringOptions struct {
	Value    string
	NilValue string
	Handler  func(string) string
}

func (p *Parser) setLowestIndex(lowest int) {
	if lowest == 0 {
		p.LowestWasZero = true
		return
	}
	if p.LowestIndex > lowest {
		p.LowestIndex = lowest
	}
}

func (p *Parser) FindString(attr string, rx *regexp.Regexp, options FindStringOptions) string {
	loc := rx.FindStringSubmatchIndex(p.Name)

	if len(loc) == 0 {
		return options.NilValue
	}

	if len(loc) == 4 && p.MatchedRange(loc[2], loc[3]) {
		return options.NilValue
	} else if len(loc) == 2 && p.MatchedRange(loc[0], loc[1]) {
		return options.NilValue
	}

	p.setLowestIndex(loc[0])

	var name string
	if len(loc) == 4 {
		p.AddMatchedIndex(attr, []int{loc[2], loc[3]})
		name = p.Name[loc[2]:loc[3]]
	} else {
		p.AddMatchedIndex(attr, []int{loc[0], loc[1]})
		name = p.Name[loc[0]:loc[1]]
	}

	if options.Value != "" {
		return options.Value
	}

	if options.Handler != nil {
		return options.Handler(name)
	}

	return name
}

type FindNumberOptions struct {
	Value    int
	NilValue int
	Handler  func(int) int
	Cleaner  func(string) string
}

func (p *Parser) FindNumber(attr string, rx *regexp.Regexp, options FindNumberOptions) int {
	loc := rx.FindStringSubmatchIndex(p.Name)

	if len(loc) == 0 {
		return options.NilValue
	}

	if len(loc) == 4 && p.MatchedRange(loc[2], loc[3]) {
		return options.NilValue
	} else if len(loc) == 2 && p.MatchedRange(loc[0], loc[1]) {
		return options.NilValue
	}

	p.setLowestIndex(loc[0])

	var name string
	if len(loc) == 4 {
		name = p.Name[loc[2]:loc[3]]
		p.AddMatchedIndex(attr, []int{loc[2], loc[3]})
	} else {
		name = p.Name[loc[0]:loc[1]]
		p.AddMatchedIndex(attr, []int{loc[0], loc[1]})
	}

	if options.Value != 0 {
		return options.Value
	}

	if options.Cleaner != nil {
		name = options.Cleaner(name)
	}

	number, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println("FindNumber:", err)
		return 0
	}

	if options.Handler != nil {
		return options.Handler(number)
	}

	return number
}
