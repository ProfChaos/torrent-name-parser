package torrentparser

import (
	"regexp"
	"strings"
)

var (
	languageGeneral *regexp.Regexp
)

func init() {
	languageGeneral = regexp.MustCompile(`(?i)\b(?:RUS|NL|FLEMISH|GERMAN|DUBBED|FR(?:ENCH)?|Truefrench|VF(?:[FI])|VOST(?:(?:F(?:R)?)|A)?|SUBFRENCH|MULTi(?:Lang|-VF2)?|ITA(?:LIAN)?|(?:iTALiAN)|Eng)\b`)
}

func (p *parser) GetLanguages() []string {
	return p.FindStrings("languages", languageGeneral, FindStringsOptions{Handler: func(strs []string) []string {
		lowerStrs := make([]string, len(strs))
		for i, str := range strs {
			lowerStrs[i] = strings.ToLower(str)
		}

		return lowerStrs
	}})
}
