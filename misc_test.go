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
