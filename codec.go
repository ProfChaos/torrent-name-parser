package torrentparser

import (
	"regexp"
	"strings"
)

var (
	codecGeneral = regexp.MustCompile(`(?i)\b(?:vix|mpeg2|divx|xvid|[xh][-. ]?26[45]|avc|hevc)\b`)

	// Regex for replacing
	codecRemove = regexp.MustCompile(`(?i)[ .-]`)
)

func (p *parser) GetCodec() string {
	return p.FindString("codec", codecGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(codecRemove.ReplaceAllString(str, ""))
	}})
}
