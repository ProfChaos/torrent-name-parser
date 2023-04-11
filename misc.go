package torrentparser

import (
	"regexp"
)

var (
	remasterGeneral,
	regionGeneral,
	hardcodedGeneral,
	containerGeneral,
	hdrGeneral,
	unratedGeneral *regexp.Regexp
)

func init() {
	unratedGeneral = regexp.MustCompile(`(?i)\bunrated|uncensored\b`)
	remasterGeneral = regexp.MustCompile(`(?i)\bRemaster(?:ed)?\b`)
	hardcodedGeneral = regexp.MustCompile(`(?i)\bHC|HARDCODED\b`)
	regionGeneral = regexp.MustCompile(`(?i)dvd(R[0-9])`)
	containerGeneral = regexp.MustCompile(`(?i)\.(MKV|AVI|MP4)$`)
	hdrGeneral = regexp.MustCompile("(?i)hdr")
}

func (p *parser) GetUnrated() bool {
	return p.FindBoolean("unrated", unratedGeneral)
}

func (p *parser) GetRemastered() bool {
	return p.FindBoolean("remaster", remasterGeneral)
}

func (p *parser) GetHardcoded() bool {
	return p.FindBoolean("hardcoded", hardcodedGeneral)
}

func (p *parser) GetRegion() string {
	return p.FindString("region", regionGeneral, FindStringOptions{})
}

func (p *parser) GetContainer() string {
	return p.FindString("container", containerGeneral, FindStringOptions{})
}

func (p *parser) GetHdr() bool {
	return p.FindBoolean("hdr", hdrGeneral)
}
