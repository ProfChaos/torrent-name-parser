package torrentparser

import (
	"regexp"
	"strings"
)

var (
	resolutionX       = regexp.MustCompile(`(?i)[0-9]{3,4}x([0-9]{3,4})`)
	resolution4k      = regexp.MustCompile(`(?i)\b4k\b`)
	resolution8k      = regexp.MustCompile(`(?i)\b8k\b`)
	resolutionGeneral = regexp.MustCompile(`(?i)[0-9]{3,4}[pi]`)
)

type Resolution string

const (
	Resolution480p    Resolution = "480p"
	Resolution576p    Resolution = "576p"
	Resolution720p    Resolution = "720p"
	Resolution1080i   Resolution = "1080i"
	Resolution1080p   Resolution = "1080p"
	Resolution4k      Resolution = "2160p"
	Resolution8k      Resolution = "4320p"
	ResolutionUnknown Resolution = ""
)

func (r Resolution) Verify() bool {
	switch r {
	case Resolution480p, Resolution576p, Resolution720p, Resolution1080i, Resolution1080p, Resolution4k, Resolution8k:
		return true
	default:
		return false
	}
}

func (p *parser) GetResolution() Resolution {
	resolution := p.FindString("resolution", resolutionX, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str) + "p"
	}})
	if resolution != "" {
		return Resolution(resolution)
	}
	resolution = p.FindString("resolution", resolutionGeneral, FindStringOptions{Handler: func(str string) string {
		return strings.ToLower(str)
	}})
	if resolution != "" {
		return Resolution(resolution)
	}
	resolution = p.FindString("resolution", resolution4k, FindStringOptions{Value: string(Resolution4k)})
	if resolution != "" {
		return Resolution(resolution)
	}
	return Resolution(p.FindString("resolution", resolution8k, FindStringOptions{Value: string(Resolution8k)}))

}
