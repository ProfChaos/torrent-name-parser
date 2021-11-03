package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	codecRemove,
	codecGeneral *regexp.Regexp
)

func init() {
	var err error
	codecGeneral, err = regexp.Compile("(?i)dvix|mpeg2|divx|xvid|[xh][-. ]?26[45]|avc|hevc")
	if err != nil {
		log.Fatalln(err)
	}
	codecRemove, err = regexp.Compile("(?i)[ .-]")
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetCodec() string {
	return p.FindString("codec", codecGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(codecRemove.ReplaceAllString(str, ""))
	}})
}
