package torrentparser

import (
	"regexp"
	"strings"
)

var (
	resolutionX       = regexp.MustCompile(`(?i)[0-9]{3,4}x([0-9]{3,4})`)
	resolution4k      = regexp.MustCompile(`(?i)\b(4k|2160p)\b`)
	resolution8k      = regexp.MustCompile(`(?i)\b(8k|4320p)\b`)
	resolutionGeneral = regexp.MustCompile(`(?i)[0-9]{3,4}[pi]`)
)

func (p *parser) GetResolution() string {
	resolution := p.FindString("resolution", resolutionX, FindStringOptions{})
	if resolution != "" {
		return resolution
	}
	resolution = p.FindString("resolution", resolution4k, FindStringOptions{Value: "4k"})
	if resolution != "" {
		return resolution
	}
	resolution = p.FindString("resolution", resolution8k, FindStringOptions{Value: "8k"})
	if resolution != "" {
		return resolution
	}
	return p.FindString("resolution", resolutionGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
}
