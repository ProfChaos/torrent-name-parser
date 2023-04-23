package torrentparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_GetTitle(t *testing.T) {

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "[HorribleSubs] Boruto - Naruto Next Generations - 85 [720p].mkv",
			want: []string{"Boruto - Naruto Next Generations"},
		},
		{
			name: "American.Dad.S17E17.720p.WEBRip.x264-BAE.mkv",
			want: []string{"American Dad"},
		},
		{
			name: "En.Affare.AKA.An.Affair.2018.1080p.BluRay.x264-HANDJOB.mkv",
			want: []string{"En Affare", "An Affair"},
		},
		{
			name: "Utøya 22. juli AKA U - 22 july [2018].mp4",
			want: []string{"Utøya 22 juli", "U - 22 july"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			titles := []string{p.Title}
			if p.AlternativeTitle != "" {
				titles = append(titles, p.AlternativeTitle)
			}

			assert.Equal(t, tt.want, titles)
		})
	}
}
