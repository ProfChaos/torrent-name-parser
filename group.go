package torrentparser

import (
	"log"
	"regexp"
	"strings"
)

var (
	groupDZon3,
	groupStart,
	groupBracketStart,
	groupEnd *regexp.Regexp
)

func init() {
	var err error
	groupDZon3, err = regexp.Compile("(?i)D-Z0N3")
	if err != nil {
		log.Fatalln(err)
	}
	groupEnd, err = regexp.Compile(`(?i)- ?([^\-. ]+)(?:\.\w+)?$`)
	if err != nil {
		log.Fatalln(err)
	}
	groupStart, err = regexp.Compile(`(?i)^(\w+)-`)
	if err != nil {
		log.Fatalln(err)
	}
	groupBracketStart, err = regexp.Compile(`(?i)^(\[([^\]]+)\])`)
	if err != nil {
		log.Fatalln(err)
	}

}

func (p *Parser) GetGroup() string {
	group := p.FindString("group", groupDZon3, FindStringOptions{})
	if group != "" {
		return group
	}
	group = p.FindString("group", groupEnd, FindStringOptions{Handler: func(str string) string {
		name := strings.TrimPrefix(str, "-")
		name = strings.TrimPrefix(name, " ")
		return name
	}})
	if group != "" {
		return group
	}
	group = p.FindString("group", groupStart, FindStringOptions{})
	if group != "" {
		return group
	}
	group = p.FindString("group", groupBracketStart, FindStringOptions{Handler: func(str string) string {
		name := strings.TrimPrefix(str, "[")
		name = strings.TrimSuffix(name, "]")
		return name
	}})
	return group
}
