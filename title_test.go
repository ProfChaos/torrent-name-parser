package torrentparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser_GetTitle(t *testing.T) {

	tests := []struct {
		name             string
		title            string
		alternativeTitle string
	}{
		{
			name:  "[HorribleSubs] Boruto - Naruto Next Generations - 85 [720p].mkv",
			title: "Boruto - Naruto Next Generations",
		},
		{
			name:  "American.Dad.S17E17.720p.WEBRip.x264-BAE.mkv",
			title: "American Dad",
		},
		{
			name:             "En.Affare.AKA.An.Affair.2018.1080p.BluRay.x264-HANDJOB.mkv",
			title:            "En Affare",
			alternativeTitle: "An Affair",
		},
		{
			name:             "Utøya 22. juli AKA U - 22 july [2018].mp4",
			title:            "Utøya 22 juli",
			alternativeTitle: "U - 22 july",
		},
		{
			name:  "The Wizard of Oz (1939) (2160p BluRay AI x265 HEVC 10bit DDP 5.1 Joy) [UTR]",
			title: "The Wizard of Oz",
		},
		{
			name:  "Black.Panther.Wakanda.Forever.2022.2160p.UHD.BluRay.x265-SURCODE.mkv",
			title: "Black Panther Wakanda Forever",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if !cmp.Equal(tt.title, p.Title) {
				t.Errorf("diff -want +got: %s", cmp.Diff(tt.title, p.Title))
			}
			if !cmp.Equal(tt.alternativeTitle, p.AlternativeTitle) {
				t.Errorf("diff -want +got: %s", cmp.Diff(tt.alternativeTitle, p.AlternativeTitle))
			}
		})
	}
}
