package torrentparser

import (
	"regexp"
	"strings"
)

var (
	audioGeneral,
	audioAc3,
	audioDd51,
	audioAac *regexp.Regexp
)

func init() {
	audioGeneral = regexp.MustCompile("(?i)MD|MP3|mp3|FLAC|Atmos|DTS(?:-HD)?|TrueHD|Dual[- ]Audio")
	audioAc3 = regexp.MustCompile("(?i)AC-?3(?:.5.1)?")
	audioDd51 = regexp.MustCompile("(?i)DD5[. ]?1")
	audioAac = regexp.MustCompile("(?i)AAC(?:[. ]?2[. ]0)?")
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
	audio = p.FindString("audio", audioDd51, FindStringOptions{Value: "dd5.1"})
	if audio != "" {
		return audio
	}
	return p.FindString("audio", audioAac, FindStringOptions{Value: "aac"})
}
