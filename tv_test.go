package torrentparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser_GetSeason(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "The Simpsons S28E21 720p HDTV x264-AVS",
			want: 28,
		},
		{
			name: "breaking.bad.s01e01.720p.bluray.x264-reward",
			want: 1,
		},
		{
			name: "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni",
			want: 1,
		},
		{
			name: "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV",
			want: 8,
		},
		{
			name: " Orange Is The New Black Season 5 Episodes 1-10 INCOMPLETE (LEAKED)",
			want: 5,
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			want: 3,
		},
		{
			name: "FakeShow.S01.1080p.BluRay.REMUX.AVC.DTS-HD.MA.5.1-NOGRP",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Season; got != tt.want {
				t.Errorf("Parser.GetSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetSeasonsNoEpisodes(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Hero Mask S1 + S2 + Extras [npz][US BD REMUX, 1080p]",
			want: []int{1, 2},
		},
		{
			name: "Attack.on.Titan.S01.S02.S03.1080p.Blu-Ray.Remux.Dual-Audio.TrueHD",
			want: []int{1, 2, 3},
		},
		{
			name: "New.Game.S02.CR.WEB-DL.AAC2.0.x264-HorribleSubs",
			want: []int{2},
		},
		{
			name: "Black Hollywood: 'They've Gotta Have Us' S01 complete (BBC, 2018) (1280x720p HD, 50fps, soft Eng subs)",
			want: []int{1},
		},
		{
			name: "New.Girl.S07.Season.7.Complete.1080p.NF.WEB.x264-maximersk [mrsktv]",
			want: []int{7},
		},
		{
			name: "Justified - Season 1 to 6 - Mp4 x264 AC3 1080p",
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "The Simpsons - Complete Seasons S01 to S28 (1080p, 720p, DVDRip)",
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28},
		},
		{
			name: "Friends.Complete.Series.S01-S10.720p.BluRay.2CH.x265.HEVC-PSA",
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Homeland.Season.1-4.Complete.720p.HDTV.X264-MRSK",
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Black Lagoon (Seasons 1-2 + OVAs) (BD 1080p)(HEVC x265 10bit)(Dual-Audio)(Eng-Subs)-Judas[TGx]",
			want: []int{1, 2},
		},
		{
			name: "Game Of Thrones Complete Season 1,2,3,4,5,6,7 406p mkv + Subs",
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Futurama Season 1 2 3 4 5 6 7 + 4 Movies - threesixtyp",
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Breaking Bad Complete Season 1 , 2 , 3, 4 ,5 ,1080p HEVC",
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "True Blood Season 1, 2, 3, 4, 5 & 6 + Extras BDRip TSV",
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "The Simpsons Season 20 21 22 23 24 25 26 27 - threesixtyp",
			want: []int{20, 21, 22, 23, 24, 25, 26, 27},
		},
		{
			name: "Perdidos: Lost: Castellano: Temporadas 1 2 3 4 5 6 (Serie Com",
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "The Boondocks Season 1, 2 & 3",
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if !cmp.Equal(tt.want, p.Seasons) {
				t.Errorf("diff -want +got: %s", cmp.Diff(tt.want, p.Seasons))
			}
			if got := p.Episode; got != 0 {
				t.Errorf("Parser.GetSeasonsNoEpisodes() = %v, want %v", got, 0)
			}
		})
	}
}

func TestParser_GetEpisode(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "The Simpsons S28E21 720p HDTV x264-AVS",
			want: 21,
		},

		{
			name: "breaking.bad.s01e01.720p.bluray.x264-reward",
			want: 1,
		},

		{
			name: "Dragon Ball Super S01 E23 French 1080p HDTV H264-Kesni",
			want: 23,
		},

		{
			name: "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV",
			want: 11,
		},

		{
			name: " Anubis saison 01 episode 38 tvrip FR",
			want: 38,
		},

		{
			name: " Le Monde Incroyable de Gumball - Saison 5 Ep 14 - L'extérieur",
			want: 14,
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			want: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := ParseName(tt.name)
			if got := p.Episode; got != tt.want {
				t.Errorf("Parser.GetEpisode() = %v, want %v", got, tt.want)
			}
		})
	}
}
