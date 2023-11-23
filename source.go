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
	sourceGeneral = regexp.MustCompile(`(?i)\b(?:(?:HD-?)?CAM|(?:DVD|VHS)scr|TC|HDTV|R5|Blu-?ray|WEB-?(?:DL)?|(?:BR|BD|DVD|PPV|WEB|HD|(?:HD-?)?TV)-?Rip|(?:DL|WEB|BD|BR|BDRE|RE)MUX)\b`)
	sourceTelesync = regexp.MustCompile(`(?i)\b(?:HD-?)?T(?:ELE)?S(?:YNC)?\b`)
	sourceDvd = regexp.MustCompile(`(?i)(DVD)(?:R[0-9])?`)
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
