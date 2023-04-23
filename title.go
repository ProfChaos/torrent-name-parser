package torrentparser

import (
	"regexp"
	"strings"
)

var (
	akaTitle *regexp.Regexp
)

func init() {
	akaTitle = regexp.MustCompile(`(?i)(?:aka)(.+)`)
}

func (p *parser) GetTitles() []string {

	title := ""
	if p.LowestWasZero {
		startIndex := 0
		for _, r := range p.MatchedIndicies {
			if r.Start == 0 {
				startIndex = r.End
			}
		}
		stopIndex := p.LowestIndex
		for _, r := range p.MatchedIndicies {
			if r.Start > startIndex && r.Start < stopIndex {
				stopIndex = r.Start
			}
		}
		title = p.Name[startIndex:stopIndex]
	} else {
		title = p.Name[0:p.LowestIndex]
	}

	title = strings.ReplaceAll(title, ".", " ")
	title = strings.ReplaceAll(title, "  ", " ")
	title = strings.Trim(title, " ")

	loc := akaTitle.FindStringSubmatchIndex(title)
	if len(loc) > 3 {
		return []string{
			strings.Trim(title[0:loc[0]], " "),
			strings.Trim(title[loc[2]:], " "),
		}
	}

	return []string{title}
}
