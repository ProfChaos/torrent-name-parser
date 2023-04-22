package torrentparser

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

var (
	removeNonDigits *regexp.Regexp
)

func init() {
	removeNonDigits = regexp.MustCompile("[^0-9]+")
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
	Title            string      `json:"title"`
	AlternativeTitle string      `json:"alternativeTitle"`
	ContentType      ContentType `json:"contentType"`
	Year             int         `json:"year"`
	Resolution       string      `json:"resolution"`
	Extended         bool        `json:"extended"`
	Unrated          bool        `json:"unrated"`
	Proper           bool        `json:"proper"`
	Repack           bool        `json:"repack"`
	Convert          bool        `json:"convert"`
	Hardcoded        bool        `json:"hardcoded"`
	Retail           bool        `json:"retail"`
	Remastered       bool        `json:"remastered"`
	Region           string      `json:"region"`
	Container        string      `json:"container"`
	Source           string      `json:"source"`
	Codec            string      `json:"codec"`
	Audio            string      `json:"audio"`
	Group            string      `json:"group"`
	Season           int         `json:"season"`
	Episode          int         `json:"episode"`
	Languages        []string    `json:"languages"`
	Hdr              bool        `json:"hdr"`
	ColorDepth       string      `json:"colorDepth"`
	Date             string      `json:"date"`
}

func (t *Torrent) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		return json.Unmarshal([]byte(v), t)
	case []byte:
		return json.Unmarshal(v, t)
	}
	return fmt.Errorf("unsupported type: %T", value)
}

func (t Torrent) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return string(b), err
}

type parser struct {
	Name            string
	MatchedIndicies map[string]index
	LowestIndex     int
	LowestWasZero   bool
}

type index struct {
	Name  string
	Start int
	End   int
}

func ParseName(name string) (Torrent, error) {
	p := &parser{Name: name, MatchedIndicies: map[string]index{}, LowestIndex: len(name)}
	return p.Parse()
}

func (p *parser) Parse() (Torrent, error) {
	torrent := Torrent{
		Season: -1,
	}

	torrent.Container = p.GetContainer()
	torrent.Resolution = p.GetResolution()
	torrent.Date = p.GetDate()
	torrent.Year = p.GetYear()
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
	torrent.Languages = p.GetLanguages()

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

func (p *parser) MatchedRange(start, end int) bool {
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

func (p *parser) AddMatchedIndex(attr string, loc []int) {
	p.MatchedIndicies[attr] = index{Name: attr, Start: loc[0], End: loc[1]}
}

func (p *parser) FindBoolean(attr string, rx *regexp.Regexp) bool {
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

type FindStringsOptions struct {
	NilValue []string
	Handler  func([]string) []string
}

func (p *parser) setLowestIndex(lowest int) {
	if lowest == 0 {
		p.LowestWasZero = true
		return
	}
	if p.LowestIndex > lowest {
		p.LowestIndex = lowest
	}
}

func (p *parser) FindString(attr string, rx *regexp.Regexp, options FindStringOptions) string {
	loc := rx.FindStringSubmatchIndex(p.Name)

	name, returnNil := p.shouldReturnNil(attr, loc)
	if returnNil {
		return options.NilValue
	}

	if options.Value != "" {
		return options.Value
	}

	if options.Handler != nil {
		return options.Handler(name)
	}

	return name
}

func (p *parser) FindStrings(attr string, rx *regexp.Regexp, options FindStringsOptions) []string {
	locs := rx.FindAllStringSubmatchIndex(p.Name, -1)

	vals, returnNil := p.shouldAllReturnNil(attr, locs)
	if returnNil {
		return options.NilValue
	}

	if options.Handler != nil {
		return options.Handler(vals)
	}

	return vals
}

type FindNumberOptions struct {
	Value    int
	NilValue int
	Handler  func(int) int
	Cleaner  func(string) string
}

func (p *parser) FindLastNumber(attr string, rx *regexp.Regexp, options FindNumberOptions) int {
	locs := rx.FindAllStringSubmatchIndex(p.Name, -1)

	if len(locs) == 0 {
		return options.NilValue
	}

	return p.parseNumber(attr, locs[len(locs)-1], options)
}

func (p *parser) FindNumber(attr string, rx *regexp.Regexp, options FindNumberOptions) int {
	loc := rx.FindStringSubmatchIndex(p.Name)

	return p.parseNumber(attr, loc, options)
}

func (p *parser) parseNumber(attr string, loc []int, options FindNumberOptions) int {
	name, returnNil := p.shouldReturnNil(attr, loc)
	if returnNil {
		return options.NilValue
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
		return options.NilValue
	}

	if options.Handler != nil {
		return options.Handler(number)
	}

	return number
}

func (p *parser) shouldReturnNil(name string, loc []int) (string, bool) {
	if len(loc) == 0 {
		return "", true
	}

	if len(loc) == 4 && p.MatchedRange(loc[2], loc[3]) {
		return "", true
	} else if len(loc) == 2 && p.MatchedRange(loc[0], loc[1]) {
		return "", true
	}

	p.setLowestIndex(loc[0])

	var match string
	if len(loc) == 4 {
		match = p.Name[loc[2]:loc[3]]
		p.AddMatchedIndex(name, []int{loc[2], loc[3]})
	} else {
		match = p.Name[loc[0]:loc[1]]
		p.AddMatchedIndex(name, []int{loc[0], loc[1]})
	}

	return match, false
}

func (p *parser) shouldAllReturnNil(name string, locs [][]int) ([]string, bool) {
	if len(locs) == 0 {
		return nil, true
	}

	matches := make([]string, 0)
	for _, loc := range locs {
		match, returnNil := p.shouldReturnNil(name, loc)
		if returnNil {
			return nil, true
		}
		matches = append(matches, match)
	}

	return matches, false
}
