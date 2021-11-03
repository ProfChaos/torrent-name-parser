package torrentparser

import (
	"log"
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
	var err error
	unratedGeneral, err = regexp.Compile(`(?i)\bunrated|uncensored\b`)
	if err != nil {
		log.Fatalln(err)
	}
	remasterGeneral, err = regexp.Compile(`(?i)\bRemaster(?:ed)?\b`)
	if err != nil {
		log.Fatalln(err)
	}
	hardcodedGeneral, err = regexp.Compile(`(?i)\bHC|HARDCODED\b`)
	if err != nil {
		log.Fatalln(err)
	}
	regionGeneral, err = regexp.Compile(`(?i)dvd(R[0-9])`)
	if err != nil {
		log.Fatalln(err)
	}
	containerGeneral, err = regexp.Compile(`(?i)\.(MKV|AVI|MP4)$`)
	if err != nil {
		log.Fatalln(err)
	}

	hdrGeneral, err = regexp.Compile("(?i)hdr")
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Parser) GetUnrated() bool {
	return p.FindBoolean("unrated", unratedGeneral)
}

func (p *Parser) GetRemastered() bool {
	return p.FindBoolean("remaster", remasterGeneral)
}

func (p *Parser) GetHardcoded() bool {
	return p.FindBoolean("hardcoded", hardcodedGeneral)
}

func (p *Parser) GetRegion() string {
	return p.FindString("region", regionGeneral, FindStringOptions{})
}

func (p *Parser) GetContainer() string {
	return p.FindString("container", containerGeneral, FindStringOptions{})
}

func (p *Parser) GetHdr() bool {
	return p.FindBoolean("hdr", hdrGeneral)
}
