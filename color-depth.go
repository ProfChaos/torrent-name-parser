package torrentparser

import (
	"regexp"
)

var (
	tenBit   = regexp.MustCompile(`(?i)10-?bit`)
	eightBit = regexp.MustCompile(`(?i)8-?bit`)
)

func (p *parser) GetColorDepth() string {
	colorDepth := p.FindString("color-depth", tenBit, FindStringOptions{Value: "10-bit"})
	if colorDepth != "" {
		return colorDepth
	}
	return p.FindString("color-depth", eightBit, FindStringOptions{Value: "8-bit"})
}
