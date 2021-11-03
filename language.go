package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	languageGeneral *regexp.Regexp
)

func init() {
	var err error
	languageGeneral, err = regexp.Compile(`(?i)\bRUS|NL|FLEMISH|GERMAN|DUBBED|FR(?:ENCH)?|Truefrench|VF(?:[FI])|VOST(?:(?:F(?:R)?)|A)?|SUBFRENCH|MULTi(?:Lang|-VF2)?|ITA(?:LIAN)?|(?:iTALiAN)\b`)
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetLanguage() string {
	return p.FindString("language", languageGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
}
