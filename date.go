package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	dateDate,
	dateYear *regexp.Regexp
)

func init() {
	var err error
	dateDate, err = regexp.Compile("(?i)[0-9]{4}.[0-9]{2}.[0-9]{2}")
	if err != nil {
		log.Fatalln(err)
	}
	dateYear, err = regexp.Compile(`(?i)(?:\s|_|\.|\(|\[)(\d{4})(?:\s|_|\.|\)|\])`)
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetDate() string {
	return p.FindString("date", dateDate, FindStringOptions{Handler: func(s string) string {
		return strings.ReplaceAll(s, ".", "-")
	}})
}

func (p *Parser) GetYear() int {
	return p.FindNumber("year", dateYear, FindNumberOptions{
		Cleaner: func(str string) string {
			return removeNonDigits.ReplaceAllString(str, "")
		},
	})
}
