package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	resolutionX,
	resolution4k,
	resolution8k,
	resolutionGeneral *regexp.Regexp
)

func init() {
	var err error
	resolutionX, err = regexp.Compile("(?i)[0-9]{3,4}x([0-9]{3,4})")
	if err != nil {
		log.Fatalln(err)
	}
	resolution4k, err = regexp.Compile("(?i)(4k|2160p)")
	if err != nil {
		log.Fatalln(err)
	}
	resolution8k, err = regexp.Compile("(?i)(8k|4320p)")
	if err != nil {
		log.Fatalln(err)
	}
	resolutionGeneral, err = regexp.Compile("(?i)[0-9]{3,4}[pi]")
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetResolution() string {
	resolution := p.FindString("resolution", resolutionX, FindStringOptions{})
	if resolution != "" {
		return resolution
	}
	resolution = p.FindString("resolution", resolution4k, FindStringOptions{Value: "4k"})
	if resolution != "" {
		return resolution
	}
	resolution = p.FindString("resolution", resolution8k, FindStringOptions{Value: "8k"})
	if resolution != "" {
		return resolution
	}
	return p.FindString("resolution", resolutionGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
}
