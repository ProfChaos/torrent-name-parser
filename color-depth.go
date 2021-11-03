package torrentparser

import (
	"log"
	"regexp"
)

var (
	tenBit,
	eightBit *regexp.Regexp
)

func init() {
	var err error
	tenBit, err = regexp.Compile("(?i)10-?bit")
	if err != nil {
		log.Fatalln(err)
	}
	eightBit, err = regexp.Compile("(?i)8-?bit")
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetColorDepth() string {
	colorDepth := p.FindString("color-depth", tenBit, FindStringOptions{Value: "10-bit"})
	if colorDepth != "" {
		return colorDepth
	}
	return p.FindString("color-depth", eightBit, FindStringOptions{Value: "8-bit"})
}
