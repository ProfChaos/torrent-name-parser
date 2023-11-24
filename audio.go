package torrentparser

import (
	"regexp"
	"strings"
)

var (
	atmosRx,
	audioGeneral,
	audioAc3,
	audioDolbyDigital,
	audioAac *regexp.Regexp
)

func init() {
	audioGeneral = regexp.MustCompile(`(?i)\b(?:MD|MP3|mp3|FLAC|DTS(?:-HD)?|TrueHD|Dual[- ]Audio)\b`)
	audioAc3 = regexp.MustCompile(`(?i)\bAC-?3(?:.5.1)?\b`)
	audioDolbyDigital = regexp.MustCompile(`(?i)\bDD(?:-EX|\+)?(?:\d[. ]\d)?(?:.+Atmos)?\b`)
	audioAac = regexp.MustCompile(`(?i)\bAAC(?:[. ]?2[. ]0)?\b`)

	// Regex for replacing
	atmosRx = regexp.MustCompile(`(?i).atmos`)
}

func (p *parser) GetAudio() string {
	audio := p.FindString("audio", audioGeneral, FindStringOptions{Handler: func(resStr string) string {
		return strings.ToLower(resStr)
	}})
	if audio != "" {
		return audio
	}
	audio = p.FindString("audio", audioAc3, FindStringOptions{Value: "ac3"})
	if audio != "" {
		return audio
	}
	audio = p.FindString("audio", audioDolbyDigital, FindStringOptions{
		Handler: func(resStr string) string {
			str := strings.ReplaceAll(resStr, " ", ".")
			return atmosRx.ReplaceAllString(str, " Atmos")
		},
	})
	if audio != "" {
		return audio
	}
	return p.FindString("audio", audioAac, FindStringOptions{Value: "aac"})
}
