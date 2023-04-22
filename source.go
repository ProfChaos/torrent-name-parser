package torrentparser

import (
	"regexp"
	"strings"
)

var (
	sourceTelesync,
	sourceDvd,
	sourceGeneral *regexp.Regexp
)

func init() {
	sourceGeneral = regexp.MustCompile(`(?i)\b(?:(?:HD-?)?CAM|HD-?Rip|HDTV|BRRip|BDRip|DVDRip|DVDscr|(?:HD-?)?TVRip|TC|PPVRip|R5|VHSSCR|Blu-?ray|WEB-?DL|WEB-?Rip|(?:DL|WEB|BD|BR|BDRE|RE)MUX)|WEB\b`)
	sourceTelesync = regexp.MustCompile("(?i)\b(?:HD-?)?T(?:ELE)?S(?:YNC)?\b")
	sourceDvd = regexp.MustCompile("(?i)(DVD)(?:R[0-9])?")
}

func (p *parser) GetSource() string {
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
