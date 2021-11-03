package torrentparser

import (
	"strings"
)

func (p *Parser) GetTitle() string {

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
	title = strings.Trim(title, " ")
	return title
}
