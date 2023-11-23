package torrentparser

import (
	"regexp"
	"strings"
)

var (
	codecRemove,
	codecGeneral *regexp.Regexp
)

func init() {
	codecGeneral = regexp.MustCompile(`(?i)dvix|mpeg2|divx|xvid|[xh][-. ]?26[45]|avc|hevc`)
	codecRemove = regexp.MustCompile(`(?i)[ .-]`)
}

func (p *parser) GetCodec() string {
	return p.FindString("codec", codecGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(codecRemove.ReplaceAllString(str, ""))
	}})
}
