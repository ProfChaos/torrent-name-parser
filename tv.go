package torrentparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	episodeGeneral,
	episodeX,
	episodeEpisode,
	episodeSeason,
	episodeAnime,
	seasonRange1,
	seasonRange2,
	seasonRange3,
	seasonList,
	seasonGeneral,
	seasonX,
	seasonSaison *regexp.Regexp
)

func init() {
	// Some of these more complex regexes are adapted from https://github.com/TheBeastLT/parse-torrent-title
	// Season ranges (ie, S01-S03) - must have two capture groups to denote the start and end of the range
	seasonRange1 = regexp.MustCompile(`(?i)(?:complete\W|(?:seasons|series)?\W|\W|^)(?:s(\d{1,2})[, +/\\&-]+)+s(\d{1,2})\b`)
	seasonRange2 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?(?:seasons?|[Сс]езони?|temporadas?)[. ]?[-:]?[. ]?[([]?(?:(?:(\d{1,2})[., /\\&-]+)+(\d{1,2})\b)[)\]]?`)
	seasonRange3 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?\bseasons?\b[. -]?S?(\d{1,2})[. -]?(?:to|thru|and|\+|:)[. -]?(?:s?)(\d{1,2})\b`) // two capture groups

	// Season list matches a substring list of seasons (ie, 1,2,3,4,5)
	seasonList = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?(?:seasons?|[Сс]езони?|temporadas?)[. ]?[-:]?[. ]?[([]?((?:\d{1,2}[., /\\&]+)+\d{1,2}\b)[)\]]?`)

	seasonGeneral = regexp.MustCompile(`(?i)[^\w]S([0-9]{1,2})(?: ?E[0-9]{1,2})?`)
	seasonSaison = regexp.MustCompile(`(?i)(?:\(?Saison|Season)[. _-]?([0-9]{1,2})`)
	seasonX = regexp.MustCompile(`(?i)[^\d]+([0-9]{1,2})x[0-9]{1,2}[^\d]+`)
	episodeGeneral = regexp.MustCompile("(?i)S[0-9]{1,2} ?E([0-9]{1,2})")
	episodeSeason = regexp.MustCompile(`(?i)\(Season \d+\) ([0-9]{1,3})\s`)
	episodeEpisode = regexp.MustCompile(`(?i)[ée]p(?:isode)?[. _-]?([0-9]{1,3})`)
	episodeAnime = regexp.MustCompile(`(?i)- ([0-9]{1,3}) (?:\[|\()`)
	episodeX = regexp.MustCompile("(?i)[0-9]{1,2}x([0-9]{1,2})")
}

func (p *parser) GetSeasons() []int {
	// Try identify season ranges before individually defined seasons/single seasons
	seasonList := p.FindString("seasonList", seasonList, FindStringOptions{})
	if seasonList != "" {
		seasons := potentialSeasonListToInts(seasonList)
		if seasons != nil {
			return seasons
		}
	}

	for idx, seasonRangeRX := range []*regexp.Regexp{seasonRange1, seasonRange2, seasonRange3} {
		seasons := p.FindNumbers("seasonRange", seasonRangeRX, FindNumbersOptions{})
		if seasons != nil && seasons[1] > seasons[0] {
			fmt.Printf("matched in season range on regex %d\n", idx)
			return intRange(seasons[0], seasons[1])
		}
	}

	season := p.FindNumbers("seasonGeneral", seasonGeneral, FindNumbersOptions{NilValue: nil})
	if season != nil {
		return season
	}
	season = p.FindNumbers("seasonSaison", seasonSaison, FindNumbersOptions{NilValue: nil})
	if season != nil {
		return season
	}
	return p.FindNumbers("seasonX", seasonX, FindNumbersOptions{NilValue: nil})
}

func (p *parser) GetEpisode() int {
	episode := p.FindNumber("episode", episodeGeneral, FindNumberOptions{})
	if episode != 0 {
		return episode
	}
	episode = p.FindNumber("episode", episodeEpisode, FindNumberOptions{})
	if episode != 0 {
		return episode
	}
	episode = p.FindNumber("episode", episodeSeason, FindNumberOptions{})
	if episode != 0 {
		return episode
	}
	episode = p.FindNumber("episode", episodeAnime, FindNumberOptions{})
	if episode != 0 {
		return episode
	}
	return p.FindNumber("episode", episodeX, FindNumberOptions{})
}

// potentialSeasonListToInts attempts to parse a season list separated by an unknown
// delimiter into a slice of ints.
// All expected delimiters are replaced with a pipe, then the string is split on the pipe
// to normalize inconsistent use of delimiters (ie, 1, 2 & 3).
func potentialSeasonListToInts(l string) []int {
	r := strings.NewReplacer(
		",", "|",
		".", "|",
		" ", "|",
		"/", "|",
		"\\", "|",
		"&", "|",
	)

	seasonParts := strings.Split(r.Replace(l), "|")
	seasons := make([]int, 0)
	for _, seasonPart := range seasonParts {
		if seasonPart == "" {
			continue
		}
		season, err := strconv.Atoi(seasonPart)
		if err != nil {
			return nil
		}
		seasons = append(seasons, season)
	}
	return seasons
}

// intRange returns a slice of integers from s to e inclusive.
func intRange(s, e int) []int {
	var r []int
	for i := s; i <= e; i++ {
		r = append(r, i)
	}
	return r
}
