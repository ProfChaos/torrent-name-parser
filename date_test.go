package torrentparser

import (
	"testing"
)

func TestParser_GetDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Stephen.Colbert.2023.02.15.Jim.Gaffigan.1080p.WEB.H264-JEBAITED.mkv",
			want: "2023-02-15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Date; got != tt.want {
				t.Errorf("Parser.GetDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetYear(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO",
			want: 2014,
		}, {
			name: "Hercules (2014) 1080p BrRip H264 - YIFY",
			want: 2014,
		}, {
			name: "One Shot [2014] DVDRip XViD-ViCKY",
			want: 2014,
		}, {
			name: "2012 2009 1080p BluRay x264 REPACK-METiS",
			want: 2009,
		}, {
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Year; got != tt.want {
				t.Errorf("Parser.GetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
