package torrentparser

import "testing"

func TestParser_GetAudio(t *testing.T) {
	type fields struct {
		Name            string
		MatchedIndicies map[string]Index
		LowestIndex     int
		LowestWasZero   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Nocturnal Animals 2016 VFF 1080p BluRay DTS HEVC-HD2",
			fields: fields{
				Name: "Nocturnal Animals 2016 VFF 1080p BluRay DTS HEVC-HD2",
			},
			want: "dts",
		},
		{
			name: "Gold 2016 1080p BluRay DTS-HD MA 5 1 x264-HDH",
			fields: fields{
				Name: "Gold 2016 1080p BluRay DTS-HD MA 5 1 x264-HDH",
			},
			want: "dts-hd",
		},
		{
			name: "Rain Man 1988 REMASTERED 1080p BRRip x264 AAC-m2g",
			fields: fields{
				Name: "Rain Man 1988 REMASTERED 1080p BRRip x264 AAC-m2g",
			},
			want: "aac",
		},
		{
			name: "The Vet Life S02E01 Dunk-A-Doctor 1080p ANPL WEB-DL AAC2 0 H 264-RTN",
			fields: fields{
				Name: "The Vet Life S02E01 Dunk-A-Doctor 1080p ANPL WEB-DL AAC2 0 H 264-RTN",
			},
			want: "aac",
		},
		{
			name: "Jimmy Kimmel 2017 05 03 720p HDTV DD5 1 MPEG2-CTL",
			fields: fields{
				Name: "Jimmy Kimmel 2017 05 03 720p HDTV DD5 1 MPEG2-CTL",
			},
			want: "dd5.1",
		},
		{
			name: "A Dog's Purpose 2016 BDRip 720p X265 Ac3-GANJAMAN",
			fields: fields{
				Name: "A Dog's Purpose 2016 BDRip 720p X265 Ac3-GANJAMAN",
			},
			want: "ac3",
		},
		{
			name: "Retroactive 1997 BluRay 1080p AC-3 HEVC-d3g",
			fields: fields{
				Name: "Retroactive 1997 BluRay 1080p AC-3 HEVC-d3g",
			},
			want: "ac3",
		},
		{
			name: "Tempete 2016-TrueFRENCH-TVrip-H264-mp3",
			fields: fields{
				Name: "Tempete 2016-TrueFRENCH-TVrip-H264-mp3",
			},
			want: "mp3",
		},
		{
			name: "Detroit.2017.BDRip.MD.GERMAN.x264-SPECTRE",
			fields: fields{
				Name: "Detroit.2017.BDRip.MD.GERMAN.x264-SPECTRE",
			},
			want: "md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: map[string]Index{},
				LowestIndex:     tt.fields.LowestIndex,
				LowestWasZero:   tt.fields.LowestWasZero,
			}
			if got := p.GetAudio(); got != tt.want {
				t.Errorf("Parser.GetAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}
