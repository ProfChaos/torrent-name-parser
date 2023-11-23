package torrentparser

import (
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
	groupDZon3 = regexp.MustCompile(`(?i)D-Z0N3`)
	groupEnd = regexp.MustCompile(`(?i)- ?([^\-. ]+)(?:\.\w+)?$`)
	groupStart = regexp.MustCompile(`(?i)^(\w+-)`)
	groupBracketStart = regexp.MustCompile(`(?i)^(\[[^\]]+\])`)
}

func (p *parser) GetGroup() string {
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
	group = p.FindString("group", groupStart, FindStringOptions{Handler: func(str string) string {
		return strings.TrimSuffix(str, "-")
	}})
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
