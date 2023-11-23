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
	audioGeneral = regexp.MustCompile(`(?i)MD|MP3|mp3|FLAC|DTS(?:-HD)?|TrueHD|Dual[- ]Audio`)
	audioAc3 = regexp.MustCompile(`(?i)AC-?3(?:.5.1)?`)
	audioDolbyDigital = regexp.MustCompile(`(?i)DD(?:-EX|\+)?(?:\d[. ]\d)?(?:.+Atmos)?`)
	audioAac = regexp.MustCompile(`(?i)AAC(?:[. ]?2[. ]0)?`)
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
