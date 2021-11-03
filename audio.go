package torrentparser

import (
	"log"
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
	var err error
	audioGeneral, err = regexp.Compile("(?i)MD|MP3|mp3|FLAC|Atmos|DTS(?:-HD)?|TrueHD|Dual[- ]Audio")
	if err != nil {
		log.Fatalln(err)
	}
	audioAc3, err = regexp.Compile("(?i)AC-?3(?:.5.1)?")
	if err != nil {
		log.Fatalln(err)
	}
	audioDd51, err = regexp.Compile("(?i)DD5[. ]?1")
	if err != nil {
		log.Fatalln(err)
	}
	audioAac, err = regexp.Compile("(?i)AAC(?:[. ]?2[. ]0)?")
	if err != nil {
		log.Fatalln(err)
	}

}

func (p *Parser) GetAudio() string {
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
