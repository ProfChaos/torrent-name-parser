package torrentparser

import "testing"

func TestParser_GetResolution(t *testing.T) {
	tests := []struct {
		name string
		want Resolution
	}{
		{
			name: "Annabelle.2014.1080p.PROPER.HC.WEBRip.x264.AAC.2.0-RARBG",
			want: "1080p",
		},
		{
			name: "doctor_who_2005.8x12.death_in_heaven.720p_hdtv_x264-fov",
			want: "720p",
		},
		{
			name: "UFC 187 PPV 720P HDTV X264-KYR",
			want: "720p",
		},
		{
			name: "The Smurfs 2 2013 COMPLETE FULL BLURAY UHD (4K) - IPT EXCLUSIVE",
			want: "2160p",
		},
		{
			name: "A.Movie.2014.4320p.PROPER.HC.WEBRip.x264.AAC.2.0-RARBG",
			want: "4320p",
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			want: "1080p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Resolution; got != tt.want {
				t.Errorf("Parser.GetResolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
