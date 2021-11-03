package torrentparser

import (
	"testing"
)

func TestParser_GetDate(t *testing.T) {
	type fields struct {
		Name            string
		MatchedIndicies map[string]Index
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: tt.fields.MatchedIndicies,
			}
			if got := p.GetDate(); got != tt.want {
				t.Errorf("Parser.GetDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetYear(t *testing.T) {
	type fields struct {
		Name            string
		MatchedIndicies map[string]Index
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO",
			fields: fields{
				Name: "Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO",
			},
			want: 2014,
		}, {
			name: "Hercules (2014) 1080p BrRip H264 - YIFY",
			fields: fields{
				Name: "Hercules (2014) 1080p BrRip H264 - YIFY",
			},
			want: 2014,
		}, {
			name: "One Shot [2014] DVDRip XViD-ViCKY",
			fields: fields{
				Name: "One Shot [2014] DVDRip XViD-ViCKY",
			},
			want: 2014,
		}, {
			name: "2012 2009 1080p BluRay x264 REPACK-METiS",
			fields: fields{
				Name: "2012 2009 1080p BluRay x264 REPACK-METiS",
			},
			want: 2009,
		}, {
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			fields: fields{
				Name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: map[string]Index{},
			}
			if got := p.GetYear(); got != tt.want {
				t.Errorf("Parser.GetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
