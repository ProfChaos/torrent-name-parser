package torrentparser

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	removeNonDigits = regexp.MustCompile(`[^0-9]+`)
)

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
	Resolution       Resolution  `json:"resolution"`
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
	Seasons          []int       `json:"seasons"`
	Episode          int         `json:"episode"`
	Languages        []string    `json:"languages"`
	Hdr              bool        `json:"hdr"`
	HdrTypes         []string    `json:"hdrTypes"`
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

func DebugParser(name string) {
	p := &parser{Name: name, MatchedIndicies: map[string]index{}, LowestIndex: len(name)}

	start := time.Now()
	p.Parse()
	fmt.Printf("Parsing took %s\n", time.Since(start))

	for _, index := range p.MatchedIndicies {
		fmt.Printf("%s\033[34m%s\033[0m%s | \033[34m%s\033[0m\n", name[:index.Start], name[index.Start:index.End], name[index.End:], index.Name)
	}
}

func (p *parser) Parse() (Torrent, error) {
	torrent := Torrent{
		Season: -1,
	}

	torrent.Container = p.GetContainer()
	torrent.Extended = p.Extended()
	torrent.Repack = p.Repack()
	torrent.Proper = p.Proper()
	torrent.Convert = p.Convert()
	torrent.Resolution = p.GetResolution()
	if !torrent.Resolution.Verify() {
		torrent.Resolution = ResolutionUnknown
	}
	torrent.Date = p.GetDate()
	torrent.Year = p.GetYear()
	torrent.Group = p.GetGroup()
	torrent.Hardcoded = p.GetHardcoded()
	torrent.Remastered = p.GetRemastered()
	torrent.Region = p.GetRegion()
	torrent.Source = p.GetSource()
	torrent.Codec = p.GetCodec()
	torrent.Audio = p.GetAudio()
	torrent.Seasons = p.GetSeasons()
	if len(torrent.Seasons) == 1 {
		torrent.Season = torrent.Seasons[0]
	}
	torrent.Episode = p.GetEpisode()
	torrent.Unrated = p.GetUnrated()
	torrent.HdrTypes, torrent.Hdr = p.Hdr()
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
	torrent.Title, torrent.AlternativeTitle = p.GetTitles()

	if torrent.Episode > 0 || torrent.Date != "" || torrent.Season > -1 {
		torrent.ContentType = TV
	} else if torrent.Season == -1 && torrent.Episode == 0 {
		torrent.ContentType = Movie
	}

	return torrent, nil
}

func (p *parser) MatchedRange(start, end int) bool {
	for _, matchedRange := range p.MatchedIndicies {
		if start <= matchedRange.Start && end >= matchedRange.End {
			return true
		}
		if start >= matchedRange.Start && start < matchedRange.End {
			return true
		}
		if end > matchedRange.Start && end <= matchedRange.End {
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
	p.AddMatchedIndex(attr, loc)
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

	// Function expects a single result, so we take the last match
	lastName := name[len(name)-1]
	if options.Handler != nil {
		return options.Handler(lastName)
	}

	return lastName
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

	// Function expects a single result, so we take the last match
	lastName := name[len(name)-1]
	if options.Cleaner != nil {
		lastName = options.Cleaner(lastName)
	}

	number, err := strconv.Atoi(lastName)
	if err != nil {
		return options.NilValue
	}

	if options.Handler != nil {
		return options.Handler(number)
	}

	return number
}

type FindNumbersOptions struct {
	NilValue []int
}

func (p *parser) FindNumbers(attr string, rx *regexp.Regexp, options FindNumbersOptions) []int {
	locs := rx.FindAllStringSubmatchIndex(p.Name, -1)
	return p.parseNumbers(attr, locs, options)
}

func (p *parser) parseNumbers(attr string, loc [][]int, options FindNumbersOptions) []int {
	names, returnNil := p.shouldAllReturnNil(attr, loc)
	if returnNil {
		return options.NilValue
	}

	numbers := make([]int, len(names))
	for i, n := range names {
		number, err := strconv.Atoi(n)
		if err != nil {
			return options.NilValue
		}
		numbers[i] = number
	}

	return numbers
}

func (p *parser) shouldReturnNil(name string, loc []int) ([]string, bool) {
	length := len(loc)
	if length == 0 {
		return nil, true
	}

	if length > 1 && p.MatchedRange(loc[length-2], loc[length-1]) {
		return nil, true
	}

	p.setLowestIndex(loc[0])

	matches := make([]string, 0)
	// If we don't have any groups in the regex
	if length == 2 {
		matches = append(matches, p.Name[loc[0]:loc[1]])
		p.AddMatchedIndex(name, []int{loc[0], loc[1]})
	} else {
		// Support for multiple groups
		for i := 2; i < length; i += 2 {
			matches = append(matches, p.Name[loc[i]:loc[i+1]])
			p.AddMatchedIndex(name, []int{loc[i], loc[i+1]})
		}
	}

	return matches, false
}

func (p *parser) shouldAllReturnNil(name string, locs [][]int) ([]string, bool) {
	if len(locs) == 0 {
		return nil, true
	}

	matches := make([]string, 0)
	for i, loc := range locs {
		match, returnNil := p.shouldReturnNil(fmt.Sprintf("%s%d", name, i), loc)
		if returnNil {
			return nil, true
		}
		matches = append(matches, match...)
	}

	return matches, false
}
