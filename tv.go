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
	seasonGeneral,
	seasonX,
	seasonSaison *regexp.Regexp
)

func init() {
	seasonGeneral = regexp.MustCompile(`(?i)[^\w]S([0-9]{1,2})(?: ?E[0-9]{1,2})?`)
	seasonSaison = regexp.MustCompile(`(?i)(?:\(?Saison|Season)[. _-]?([0-9]{1,2})`)
	seasonX = regexp.MustCompile(`(?i)[^\d]+([0-9]{1,2})x[0-9]{1,2}[^\d]+`)
	episodeGeneral = regexp.MustCompile("(?i)S[0-9]{1,2} ?E([0-9]{1,2})")
	episodeSeason = regexp.MustCompile(`(?i)\(Season \d+\) ([0-9]{1,3})\s`)
	episodeEpisode = regexp.MustCompile(`(?i)[ée]p(?:isode)?[. _-]?([0-9]{1,3})`)
	episodeAnime = regexp.MustCompile(`(?i)- ([0-9]{1,3}) (?:\[|\()`)
	episodeX = regexp.MustCompile("(?i)[0-9]{1,2}x([0-9]{1,2})")
}

func (p *Parser) GetSeason() int {
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

func (p *Parser) GetEpisode() int {
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
