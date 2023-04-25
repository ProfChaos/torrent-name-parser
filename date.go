package torrentparser

import (
	"regexp"
	"strings"
)

var (
	dateDate,
	dateYear *regexp.Regexp
)

func init() {
	dateDate = regexp.MustCompile(`(?i)[0-9]{4}.[0-9]{2}.[0-9]{2}`)
	dateYear = regexp.MustCompile(`(?i)(?:\[|\(|\b)(\d{4})(?:\]|\)|\b)`)
}

func (p *parser) GetDate() string {
	return p.FindString("date", dateDate, FindStringOptions{Handler: func(s string) string {
		return strings.ReplaceAll(s, ".", "-")
	}})
}

func (p *parser) GetYear() int {
	return p.FindLastNumber("year", dateYear, FindNumberOptions{
		Cleaner: func(str string) string {
			return removeNonDigits.ReplaceAllString(str, "")
		},
	})
}
