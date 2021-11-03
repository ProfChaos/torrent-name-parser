package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	sourceTelesync,
	sourceDvd,
	sourceGeneral *regexp.Regexp
)

func init() {
	var err error
	sourceGeneral, err = regexp.Compile(`(?i)\b(?:HD-?)?CAM|HD-?Rip|HDTV|BRRip|BDRip|DVDRip|DVDscr|(?:HD-?)?TVRip|TC|PPVRip|R5|VHSSCR|Bluray|WEB-?DL|WEB-?Rip|(?:DL|WEB|BD|BR)MUX\b`)
	if err != nil {
		log.Fatalln(err)
	}
	sourceTelesync, err = regexp.Compile("(?i)\b(?:HD-?)?T(?:ELE)?S(?:YNC)?\b")
	if err != nil {
		log.Fatalln(err)
	}
	sourceDvd, err = regexp.Compile("(?i)(DVD)(?:R[0-9])?")
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetSource() string {
	source := p.FindString("sourceGeneral", sourceGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
	if source != "" {
		return source
	}
	source = p.FindString("sourceTelesync", sourceTelesync, FindStringOptions{Value: "telesync"})
	if source != "" {
		return source
	}
	return p.FindString("sourceDvd", sourceDvd, FindStringOptions{Value: "dvd"})
}
