package torrentparser

import "testing"

func TestParser_GetSource(t *testing.T) {
	tests := []struct {
		desc string
		name string
		want string
	}{
		{
			desc: "Not tc from the word watchmen",
			name: "Watchmen.The.Ultimate.Cut.2009.1080p.BluRay.AC3.x264-CtrlHD.mkv",
			want: "bluray",
		},
		{
			desc: "Not tc from the word pitch",
			name: "Pitch.Perfect.2.2015.REPACK.1080p.BluRay.DTS.x264-VietHD.mkv",
			want: "bluray",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Source; got != tt.want {
				t.Errorf("Parser.GetSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
