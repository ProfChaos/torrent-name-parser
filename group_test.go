package torrentparser

import "testing"

func TestParser_GetGroup(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{

		{
			name: "Nocturnal Animals 2016 VFF 1080p BluRay DTS HEVC-HD2",
			want: "HD2",
		},
		{
			name: "[HorribleSubs] Boruto - Naruto Next Generations - 85 [720p].mkv",
			want: "HorribleSubs",
		},

		{
			name: "X-Men.Apocalypse.2016.1080p.BluRay.DTS.x264.D-Z0N3.mkv",
			want: "D-Z0N3",
		},

		{
			name: "Gold 2016 1080p BluRay DTS-HD MA 5 1 x264-HDH",
			want: "HDH",
		},

		{
			name: "Hercules (2014) 1080p BrRip H264 - YIFY",
			want: "YIFY",
		},
		{
			name: "Western - L'homme qui n'a pas d'étoile-1955.Multi.DVD9",
			want: "",
		},
		{
			name: "sons.of.anarchy.s05e10.480p.BluRay.x264-GAnGSteR",
			want: "GAnGSteR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParseName(tt.name)
			if got := p.GetGroup(); got != tt.want {
				t.Errorf("Parser.GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
