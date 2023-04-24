package torrentparser

import "testing"

func TestParser_GetAudio(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Nocturnal Animals 2016 VFF 1080p BluRay DTS HEVC-HD2",
			want: "dts",
		},
		{
			name: "Gold 2016 1080p BluRay DTS-HD MA 5 1 x264-HDH",
			want: "dts-hd",
		},
		{
			name: "Rain Man 1988 REMASTERED 1080p BRRip x264 AAC-m2g",
			want: "aac",
		},
		{
			name: "The Vet Life S02E01 Dunk-A-Doctor 1080p ANPL WEB-DL AAC2 0 H 264-RTN",
			want: "aac",
		},
		{
			name: "Jimmy Kimmel 2017 05 03 720p HDTV DD5 1 MPEG2-CTL",
			want: "DD5.1",
		},
		{
			name: "A Dog's Purpose 2016 BDRip 720p X265 Ac3-GANJAMAN",
			want: "ac3",
		},
		{
			name: "Retroactive 1997 BluRay 1080p AC-3 HEVC-d3g",
			want: "ac3",
		},
		{
			name: "Detroit.2017.BDRip.MD.GERMAN.x264-SPECTRE",
			want: "md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Audio; got != tt.want {
				t.Errorf("Parser.GetAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}
