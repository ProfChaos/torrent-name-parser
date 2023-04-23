package torrentparser

import (
	"regexp"
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
	seasonRange4,
	seasonRange5,
	seasonRange6,
	seasonRange7,
	seasonGeneral,
	seasonX,
	seasonSaison *regexp.Regexp
)

func init() {
	// Season ranges (ie, S01-S03)
	//seasonRange1 = regexp.MustCompile(`(?i)(?:complete\W|seasons?\W|\W|^)((?:s\d{1,2}[., +/\\&-]+)+s\d{1,2}\b)`)
	//seasonRange2 = regexp.MustCompile(`(?i)(?:complete\W|seasons?\W|\W|^)[([]?(s\d{2,}-\d{2,}\b)[)\]]?`)
	//seasonRange3 = regexp.MustCompile(`(?i)(?:complete\W|seasons?\W|\W|^)[([]?(s[1-9]-[2-9]\b)[)\]]?`)
	//seasonRange4 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?(?:seasons?|[Сс]езони?|temporadas?)[. ]?[-:]?[. ]?[([]?((?:\d{1,2}[., /\\&]+)+\d{1,2}\b)[)\]]?`)
	//seasonRange5 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?(?:seasons|[Сс]езони?|temporadas?)[. ]?[-:]?[. ]?[([]?((?:\d{1,2}[. -]+)+[1-9]\d?\b)[)\]]?`)
	//seasonRange6 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?season[. ]?[([]?((?:\d{1,2}[. -]+)+[1-9]\d?\b)[)\]]?(!.*\.\w{2,4}$)`)
	seasonRange7 = regexp.MustCompile(`(?i)(?:(?:\bthe\W)?\bcomplete\W)?\bseasons?\b[. -]?S?(\d{1,2})[. -]?(?:to|thru|and|\+|:)[. -]?(?:s?)(\d{1,2})\b`) // two capture groups

	seasonGeneral = regexp.MustCompile(`(?i)[^\w]S([0-9]{1,2})(?: ?E[0-9]{1,2})?`)
	seasonSaison = regexp.MustCompile(`(?i)(?:\(?Saison|Season)[. _-]?([0-9]{1,2})`)
	seasonX = regexp.MustCompile(`(?i)[^\d]+([0-9]{1,2})x[0-9]{1,2}[^\d]+`)
	episodeGeneral = regexp.MustCompile("(?i)S[0-9]{1,2} ?E([0-9]{1,2})")
	episodeSeason = regexp.MustCompile(`(?i)\(Season \d+\) ([0-9]{1,3})\s`)
	episodeEpisode = regexp.MustCompile(`(?i)[ée]p(?:isode)?[. _-]?([0-9]{1,3})`)
	episodeAnime = regexp.MustCompile(`(?i)- ([0-9]{1,3}) (?:\[|\()`)
	episodeX = regexp.MustCompile("(?i)[0-9]{1,2}x([0-9]{1,2})")
}

func (p *parser) GetSeason() int {
	season := p.FindNumber("seasonGeneral", seasonGeneral, FindNumberOptions{NilValue: -1})
	if season != -1 {
		return season
	}
	season = p.FindNumber("seasonSaison", seasonSaison, FindNumberOptions{NilValue: -1})
	if season != -1 {
		return season
	}
	return p.FindNumber("seasonX", seasonX, FindNumberOptions{NilValue: -1})
}

func (p *parser) GetSeasons() []int {
	// Identify season ranges before individually defined seasons/single seasons
	//fmt.Printf("parsing: %s\n", p.Name)
	//for _, seasonRangeRX := range []*regexp.Regexp{seasonRange1, seasonRange2, seasonRange3, seasonRange4, seasonRange5, seasonRange6, seasonRange7} {
	// for _, seasonRangeRX := range []*regexp.Regexp{seasonRange7} {
	// 	seasons := p.FindNumbers("seasonRange", seasonRangeRX, FindNumberOptions{})
	// 	if seasons != nil {
	// 		fmt.Printf("Found season range: %v\n", seasons)
	// 		return seasons
	// 	}
	// 	//res := seasonRangeRX.FindAllStringSubmatchIndex(p.Name, -1)
	// }

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
