package torrentparser

import (
	"log"
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
	var err error
	seasonGeneral, err = regexp.Compile(`(?i)[^\w]S([0-9]{1,2})(?: ?E[0-9]{1,2})?`)
	if err != nil {
		log.Fatalln(err)
	}
	seasonSaison, err = regexp.Compile(`(?i)(?:\(?Saison|Season)[. _-]?([0-9]{1,2})`)
	if err != nil {
		log.Fatalln(err)
	}
	seasonX, err = regexp.Compile(`(?i)[^\d]+([0-9]{1,2})x[0-9]{1,2}[^\d]+`)
	if err != nil {
		log.Fatalln(err)
	}
	episodeGeneral, err = regexp.Compile("(?i)S[0-9]{1,2} ?E([0-9]{1,2})")
	if err != nil {
		log.Fatalln(err)
	}
	episodeSeason, err = regexp.Compile(`(?i)\(Season \d+\) ([0-9]{1,3})\s`)
	if err != nil {
		log.Fatalln(err)
	}
	episodeEpisode, err = regexp.Compile(`(?i)[ée]p(?:isode)?[. _-]?([0-9]{1,3})`)
	if err != nil {
		log.Fatalln(err)
	}
	episodeAnime, err = regexp.Compile(`(?i)- ([0-9]{1,3}) (?:\[|\()`)
	if err != nil {
		log.Fatalln(err)
	}
	episodeX, err = regexp.Compile("(?i)[0-9]{1,2}x([0-9]{1,2})")
	if err != nil {
		log.Fatalln(err)
	}
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
