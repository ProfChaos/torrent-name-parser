package torrentparser

import (
	"regexp"
	"slices"
)

var (
	remasterGeneral,
	regionGeneral,
	hardcodedGeneral,
	containerGeneral,
	hdrGeneral,
	repackGeneral,
	extendedGeneral,
	properGeneral,
	convertGeneral,
	unratedGeneral *regexp.Regexp
)

func init() {
	unratedGeneral = regexp.MustCompile(`(?i)\b(?:unrated|uncensored)\b`)
	remasterGeneral = regexp.MustCompile(`(?i)\bRemaster(?:ed)?\b`)
	hardcodedGeneral = regexp.MustCompile(`(?i)\b(?:HC|HARDCODED)\b`)
	regionGeneral = regexp.MustCompile(`(?i)\bdvd(R[0-9])\b`)
	containerGeneral = regexp.MustCompile(`(?i)\.(MKV|AVI|MP4)$`)
	hdrGeneral = regexp.MustCompile(`(?i)\b(?:hdr(?:10)?|dv)\b`)
	repackGeneral = regexp.MustCompile(`(?i)\b(?:repack|rerip)\b`)
	extendedGeneral = regexp.MustCompile(`(?i)\bextended\b`)
	properGeneral = regexp.MustCompile(`(?i)\bproper\b`)
	convertGeneral = regexp.MustCompile(`(?i)\bconvert\b`)
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

func (p *parser) Hdr() ([]string, bool) {
	isHDR := false
	return p.FindStrings("hdr", hdrGeneral, FindStringsOptions{
		Handler: func(s []string) []string {
			if len(s) > 0 {
				isHDR = true
			}
			pos := slices.Index(s, "HDR")
			if pos != -1 {
				if len(s) == 1 {
					return nil
				}

				return slices.Delete(s, pos, pos+1)
			}
			return s
		},
	}), isHDR
}

func (p *parser) Repack() bool {
	return p.FindBoolean("repack", repackGeneral)
}

func (p *parser) Extended() bool {
	return p.FindBoolean("extended", extendedGeneral)
}

func (p *parser) Proper() bool {
	return p.FindBoolean("proper", properGeneral)
}

func (p *parser) Convert() bool {
	return p.FindBoolean("convert", convertGeneral)
}
