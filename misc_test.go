package torrentparser

import "testing"

func TestParser_GetUnrated(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Identity.Thief.2013.Vostfr.UNRATED.BluRay.720p.DTS.x264-Nenuko",
			want: true,
		},
		{
			name: "Charlie.les.filles.lui.disent.merci.2007.UNCENSORED.TRUEFRENCH.DVDRiP.AC3.Libe",
			want: true,
		},
		{
			name: "Have I Got News For You S53E02 EXTENDED 720p HDTV x264-QPEL",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Unrated; got != tt.want {
				t.Errorf("Parser.GetUnrated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Repack(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Ascension.2021.REPACK.1080p.AMZN.WEB-DL.DDP5.1.H.264-WELP.mkv",
			want: true,
		},
		{
			name: "Asterix.Le.Secret.de.la.Potion.Magique.2018.FRENCH.1080p.BluRay.DTS.x264-Ulysse.mp4",
			want: false,
		},
		{
			name: "The.Tomorrow.War.2021.repack.HDR.2160p.WEB.H265-NAISU.mp4",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Repack; got != tt.want {
				t.Errorf("Parser.Repack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Extended(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Night School 2018 Extended 1080p BluRay DD+7.1 x264-DON.mkv",
			want: true,
		},
		{
			name: "Asterix.Le.Secret.de.la.Potion.Magique.2018.FRENCH.1080p.BluRay.DTS.x264-Ulysse.mp4",
			want: false,
		},
		{
			name: "Have I Got News For You S53E02 EXTENDED 720p HDTV x264-QPEL",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Extended; got != tt.want {
				t.Errorf("Parser.Extended() = %v, want %v", got, tt.want)
			}
		})
	}
}
