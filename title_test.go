package torrentparser

import "testing"

func TestParser_GetTitle(t *testing.T) {
	type fields struct {
		Name            string
		MatchedIndicies map[string]Index
		LowestWasZero   bool
		LowestIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "[HorribleSubs] Boruto - Naruto Next Generations - 85 [720p].mkv",
			fields: fields{
				Name:          "[HorribleSubs] Boruto - Naruto Next Generations - 85 [720p].mkv",
				LowestIndex:   48,
				LowestWasZero: true,
				MatchedIndicies: map[string]Index{
					"test":  {Name: "TitleTest", Start: 0, End: 15},
					"test2": {Name: "TitleTest", Start: 48, End: 56},
				},
			},
			want: "Boruto - Naruto Next Generations",
		},
		{
			name: "American.Dad.S17E17.720p.WEBRip.x264-BAE.mkv",
			fields: fields{
				Name:            "American.Dad.S17E17.720p.WEBRip.x264-BAE.mkv",
				LowestIndex:     13,
				MatchedIndicies: map[string]Index{"test": {Name: "TitleTest", Start: 13, End: 18}},
			},
			want: "American Dad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: tt.fields.MatchedIndicies,
				LowestIndex:     tt.fields.LowestIndex,
				LowestWasZero:   tt.fields.LowestWasZero,
			}
			if got := p.GetTitle(); got != tt.want {
				t.Errorf("Parser.GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
