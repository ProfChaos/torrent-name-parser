package torrentparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		Name            string
		MatchedIndicies map[string]Index
	}
	tests := []struct {
		name    string
		fields  fields
		want    Torrent
		wantErr bool
	}{
		{
			name: "sons.of.anarchy.s05e10.480p.BluRay.x264-GAnGSteR",
			fields: fields{
				Name: "sons.of.anarchy.s05e10.480p.BluRay.x264-GAnGSteR",
			},
			want: Torrent{
				Title:      "sons of anarchy",
				Resolution: "480p",
				Season:     5,
				Episode:    10,
				Source:     "bluray",
				Codec:      "x264",
				Group:      "GAnGSteR",
				TV:         true,
			},
		},
		{
			name: "Color.Of.Night.Unrated.DC.VostFR.BRrip.x264",
			fields: fields{
				Name: "Color.Of.Night.Unrated.DC.VostFR.BRrip.x264",
			},
			want: Torrent{
				Season:   -1,
				Title:    "Color Of Night",
				Unrated:  true,
				Language: "vostfr",
				Source:   "brrip",
				Codec:    "x264",
				Movie:    true,
			},
		},
		{
			name: "Da Vinci Code DVDRip",
			fields: fields{
				Name: "Da Vinci Code DVDRip",
			},
			want: Torrent{
				Season: -1,
				Title:  "Da Vinci Code",
				Source: "dvdrip",
				Movie:  true,
			},
		},
		{
			name: "Some.girls.1998.DVDRip",
			fields: fields{
				Name: "Some.girls.1998.DVDRip",
			},
			want: Torrent{
				Season: -1,
				Title:  "Some girls",
				Source: "dvdrip",
				Year:   1998,
				Movie:  true,
			},
		},
		{
			name: "Ecrit.Dans.Le.Ciel.1954.MULTI.DVDRIP.x264.AC3-gismo65",
			fields: fields{
				Name: "Ecrit.Dans.Le.Ciel.1954.MULTI.DVDRIP.x264.AC3-gismo65",
			},
			want: Torrent{
				Season:   -1,
				Title:    "Ecrit Dans Le Ciel",
				Source:   "dvdrip",
				Year:     1954,
				Language: "multi",
				Codec:    "x264",
				Audio:    "ac3",
				Group:    "gismo65",
				Movie:    true,
			},
		},
		{
			name: "2019 After The Fall Of New York 1983 REMASTERED BDRip x264-GHOULS",
			fields: fields{
				Name: "2019 After The Fall Of New York 1983 REMASTERED BDRip x264-GHOULS",
			},
			want: Torrent{
				Season:     -1,
				Title:      "2019 After The Fall Of New York",
				Source:     "bdrip",
				Remastered: true,
				Year:       1983,
				Codec:      "x264",
				Group:      "GHOULS",
				Movie:      true,
			},
		},
		{
			name: "Ghost In The Shell 2017 720p HC HDRip X264 AC3-EVO",
			fields: fields{
				Name: "Ghost In The Shell 2017 720p HC HDRip X264 AC3-EVO",
			},
			want: Torrent{
				Season:     -1,
				Title:      "Ghost In The Shell",
				Source:     "hdrip",
				Hardcoded:  true,
				Year:       2017,
				Resolution: "720p",
				Codec:      "x264",
				Audio:      "ac3",
				Group:      "EVO",
				Movie:      true,
			},
		},
		{
			name: "Rogue One 2016 1080p BluRay x264-SPARKS",
			fields: fields{
				Name: "Rogue One 2016 1080p BluRay x264-SPARKS",
			},
			want: Torrent{
				Season:     -1,
				Title:      "Rogue One",
				Source:     "bluray",
				Year:       2016,
				Resolution: "1080p",
				Codec:      "x264",
				Group:      "SPARKS",
				Movie:      true,
			},
		},
		{
			name: "Desperation 2006 Multi Pal DvdR9-TBW1973",
			fields: fields{
				Name: "Desperation 2006 Multi Pal DvdR9-TBW1973",
			},
			want: Torrent{
				Season:   -1,
				Title:    "Desperation",
				Source:   "dvd",
				Year:     2006,
				Language: "multi",
				Region:   "R9",
				Group:    "TBW1973",
				Movie:    true,
			},
		},
		{
			name:   "Maman, j'ai raté l'avion 1990 VFI 1080p BluRay DTS x265-HTG",
			fields: fields{Name: "Maman, j'ai raté l'avion 1990 VFI 1080p BluRay DTS x265-HTG"},
			want: Torrent{
				Season: -1,

				Title:      "Maman, j'ai raté l'avion",
				Source:     "bluray",
				Year:       1990,
				Audio:      "dts",
				Resolution: "1080p",
				Language:   "vfi",
				Codec:      "x265",
				Group:      "HTG",
				Movie:      true,
			},
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			fields: fields{
				Name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			},
			want: Torrent{
				Title:      "Shingeki no Kyojin",
				Audio:      "flac",
				Resolution: "1080",
				Container:  "mkv",
				Group:      "BlurayDesuYo",
				Season:     3,
				Episode:    38,
				ColorDepth: "10-bit",
				TV:         true,
			},
		},
		{
			name: "[Ohys-Raws] JoJo no Kimyou na Bouken Ougon no Kaze - 33 (BS11 1280x720 x264 AAC).mp4",
			fields: fields{
				Name: "[Ohys-Raws] JoJo no Kimyou na Bouken Ougon no Kaze - 33 (BS11 1280x720 x264 AAC).mp4",
			},
			want: Torrent{
				Season:     -1,
				Audio:      "aac",
				Codec:      "x264",
				Container:  "mp4",
				Episode:    33,
				Group:      "Ohys-Raws",
				Resolution: "720",
				Title:      "JoJo no Kimyou na Bouken Ougon no Kaze",
				TV:         true,
			},
		},
		{
			name: "[HorribleSubs] Boruto - Naruto Next Generations - 111 [720p].mkv",
			fields: fields{
				Name: "[HorribleSubs] Boruto - Naruto Next Generations - 111 [720p].mkv",
			},
			want: Torrent{
				Season:     -1,
				Container:  "mkv",
				Episode:    111,
				Group:      "HorribleSubs",
				Resolution: "720p",
				Title:      "Boruto - Naruto Next Generations",
				TV:         true,
			},
		},
		{
			name: "Marvels.Agents.of.S.H.I.E.L.D.S06E05.720p.HDTV.x264-AVS.mkv",
			fields: fields{
				Name: "Marvels.Agents.of.S.H.I.E.L.D.S06E05.720p.HDTV.x264-AVS.mkv",
			},
			want: Torrent{
				Codec:      "x264",
				Container:  "mkv",
				Episode:    5,
				Group:      "AVS",
				Resolution: "720p",
				Season:     6,
				Source:     "hdtv",
				Title:      "Marvels Agents of S H I E L D",
				TV:         true,
			},
		},
		{
			name: "stephen.colbert.2019.02.03.conan.obrien.web.x264-cookiemonster.mkv",
			fields: fields{
				Name: "stephen.colbert.2019.02.03.conan.obrien.web.x264-cookiemonster.mkv",
			},
			want: Torrent{
				Season:    -1,
				Codec:     "x264",
				Container: "mkv",
				Group:     "cookiemonster",
				Title:     "stephen colbert",
				Date:      "2019-02-03",
				TV:        true,
			},
		},
		{
			name: "Star.Wars.Episode.IX.The.Rise.of.Skywalker.2019.2160p.WEB-DL.DDP5.1.Atmos.HEVC-BLUTONiUM.mkv",
			fields: fields{
				Name: "Star.Wars.Episode.IX.The.Rise.of.Skywalker.2019.2160p.WEB-DL.DDP5.1.Atmos.HEVC-BLUTONiUM.mkv",
			},
			want: Torrent{
				Season:     -1,
				Audio:      "atmos",
				Codec:      "hevc",
				Year:       2019,
				Container:  "mkv",
				Group:      "BLUTONiUM",
				Title:      "Star Wars Episode IX The Rise of Skywalker",
				Resolution: "4k",
				Source:     "web-dl",
				Movie:      true,
			},
		},
		{
			name: "Star.Wars.Episode.7.The.Force.Awakens.2015.1080p.BluRay.DTS.x264.D-Z0N3",
			fields: fields{
				Name: "Star.Wars.Episode.7.The.Force.Awakens.2015.1080p.BluRay.DTS.x264.D-Z0N3",
			},
			want: Torrent{
				Season:     -1,
				Audio:      "dts",
				Group:      "D-Z0N3",
				Resolution: "1080p",
				Source:     "bluray",
				Codec:      "x264",
				Title:      "Star Wars Episode 7 The Force Awakens",
				Movie:      true,
				Year:       2015,
			},
		},
		{
			name: "Last.Week.Tonight.with.John.Oliver.S08E01.February.14.2021.720p.HMAX.WEB-DL.DD2.0.H.264-null.mkv",
			fields: fields{
				Name: "Last.Week.Tonight.with.John.Oliver.S08E01.February.14.2021.720p.HMAX.WEB-DL.DD2.0.H.264-null.mkv",
			},
			want: Torrent{
				Codec:      "h264",
				Container:  "mkv",
				Group:      "null",
				Resolution: "720p",
				Episode:    1,
				Season:     8,
				Source:     "web-dl",
				Title:      "Last Week Tonight with John Oliver",
				TV:         true,
				Year:       2021,
			},
		},
		{
			name: "Pirates.of.the.Caribbean.Dead.Mans.Chest.2006.2160p.DSNP.WEB-DL.DTS-HD.MA.5.1.HDR.HEVC-WATCHER.mkv",
			fields: fields{
				Name: "Pirates.of.the.Caribbean.Dead.Mans.Chest.2006.2160p.DSNP.WEB-DL.DTS-HD.MA.5.1.HDR.HEVC-WATCHER.mkv",
			},
			want: Torrent{
				Title:      "Pirates of the Caribbean Dead Mans Chest",
				Movie:      true,
				Year:       2006,
				Resolution: "4k",
				Container:  "mkv",
				Source:     "web-dl",
				Codec:      "hevc",
				Audio:      "dts-hd",
				Group:      "WATCHER",
				Season:     -1,
				Hdr:        true,
			},
		},
		{
			name: "[SubsPlease] Boku no Hero Academia - 106 (720p) [F8EFA646].mkv",
			fields: fields{
				Name: "[SubsPlease] Boku no Hero Academia - 106 (720p) [F8EFA646].mkv",
			},
			want: Torrent{
				Title:      "Boku no Hero Academia",
				TV:         true,
				Resolution: "720p",
				Container:  "mkv",
				Group:      "SubsPlease",
				Season:     -1,
				Episode:    106,
			},
		},
		{
			name: "Altered Carbon S01 2160p HDR Netflix WEBRip DD+ Atmos 5.1 x265-TrollUHD",
			fields: fields{
				Name: "Altered Carbon S01 2160p HDR Netflix WEBRip DD+ Atmos 5.1 x265-TrollUHD",
			},
			want: Torrent{
				Title:      "Altered Carbon",
				TV:         true,
				Resolution: "4k",
				Group:      "TrollUHD",
				Source:     "webrip",
				Codec:      "x265",
				Audio:      "atmos",
				Season:     1,
				Hdr:        true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				Name:            tt.fields.Name,
				MatchedIndicies: map[string]Index{},
				LowestIndex:     len(tt.fields.Name),
			}
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.Parse() = %+v, want %+v", fmt.Sprintf("%v+", got), fmt.Sprintf("%v+", tt.want))
			}
		})
	}
}
