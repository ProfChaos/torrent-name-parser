package torrentparser

import (
	"regexp"
	"strings"
)

var (
	languageGeneral *regexp.Regexp
)

func init() {
	languageGeneral = regexp.MustCompile(`(?i)\b(?:RUS|NL|FLEMISH|GERMAN|DUBBED|FR(?:ENCH)?|Truefrench|VF(?:[FI])|VOST(?:(?:F(?:R)?)|A)?|SUBFRENCH|MULTi(?:Lang|-VF2)?|ITA(?:LIAN)?|(?:iTALiAN))\b`)
}

func (p *parser) GetLanguage() string {
	return p.FindString("language", languageGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
}
