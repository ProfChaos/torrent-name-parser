package torrentparser

import (
	"testing"
)

func TestParser_GetSeason(t *testing.T) {
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
			name:   "The Simpsons S28E21 720p HDTV x264-AVS",
			fields: fields{Name: "The Simpsons S28E21 720p HDTV x264-AVS"},
			want:   28,
		},
		{
			name:   "breaking.bad.s01e01.720p.bluray.x264-reward",
			fields: fields{Name: "breaking.bad.s01e01.720p.bluray.x264-reward"},
			want:   1,
		},
		{
			name:   "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni",
			fields: fields{Name: "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni"},
			want:   1,
		},
		{
			name:   "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV",
			fields: fields{Name: "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV"},
			want:   8,
		},
		{
			name:   " Orange Is The New Black Season 5 Episodes 1-10 INCOMPLETE (LEAKED)",
			fields: fields{Name: " Orange Is The New Black Season 5 Episodes 1-10 INCOMPLETE (LEAKED)"},
			want:   5,
		},
		{
			name:   "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			fields: fields{Name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv"},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: map[string]Index{},
			}
			if got := p.GetSeason(); got != tt.want {
				t.Errorf("Parser.GetSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetEpisode(t *testing.T) {
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
			name: "The Simpsons S28E21 720p HDTV x264-AVS",
			fields: fields{
				Name: "The Simpsons S28E21 720p HDTV x264-AVS",
			},
			want: 21,
		},

		{
			name: "breaking.bad.s01e01.720p.bluray.x264-reward",
			fields: fields{
				Name: "breaking.bad.s01e01.720p.bluray.x264-reward",
			},
			want: 1,
		},

		{
			name: "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni",
			fields: fields{
				Name: "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni",
			},
			want: 23,
		},

		{
			name: "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV",
			fields: fields{
				Name: "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV",
			},
			want: 11,
		},

		{
			name: " Anubis saison 01 episode 38 tvrip FR",
			fields: fields{
				Name: " Anubis saison 01 episode 38 tvrip FR",
			},
			want: 38,
		},

		{
			name: " Le Monde Incroyable de Gumball - Saison 5 Ep 14 - L'extérieur",
			fields: fields{
				Name: " Le Monde Incroyable de Gumball - Saison 5 Ep 14 - L'extérieur",
			},
			want: 14,
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			fields: fields{
				Name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			},
			want: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: map[string]Index{},
			}
			if got := p.GetEpisode(); got != tt.want {
				t.Errorf("Parser.GetEpisode() = %v, want %v", got, tt.want)
			}
		})
	}
}
