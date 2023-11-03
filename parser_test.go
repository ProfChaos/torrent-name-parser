package torrentparser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		want    Torrent
		wantErr bool
	}{
		{
			name: "blade.runner.2049.2017.2160p.uhd.bluray.x265-terminal.mkv",
			want: Torrent{
				Title:       "blade runner 2049",
				Year:        2017,
				Resolution:  "4k",
				ContentType: Movie,
				Season:      -1,
				Group:       "terminal",
				Container:   "mkv",
				Source:      "bluray",
				Codec:       "x265",
			},
		},
		{
			name: "Wonder.Woman.1984.2020.2160p.HMAX.WEB-DL.DDP5.1.Atmos.HDR.HEVC-TOMMY.mkv",
			want: Torrent{
				Title:       "Wonder Woman 1984",
				Year:        2020,
				Resolution:  "4k",
				ContentType: Movie,
				Season:      -1,
				Group:       "TOMMY",
				Container:   "mkv",
				Hdr:         true,
				Source:      "web-dl",
				Codec:       "hevc",
				Audio:       "DDP5.1 Atmos",
			},
		},
		{
			name: "blow-action.point.2018.1080p.bluray.x264.mp4",
			want: Torrent{
				Title:       "action point",
				Resolution:  "1080p",
				Year:        2018,
				Source:      "bluray",
				Codec:       "x264",
				ContentType: Movie,
				Season:      -1,
				Group:       "blow",
				Container:   "mp4",
			},
		},
		{
			name: "Frozen.2.2019.1080p.WEB-DL.H264.AC3-EVO.mp4",
			want: Torrent{
				Title:       "Frozen 2",
				Resolution:  "1080p",
				Year:        2019,
				Source:      "web-dl",
				Codec:       "h264",
				Group:       "EVO",
				Audio:       "ac3",
				ContentType: Movie,
				Container:   "mp4",
				Season:      -1,
			},
		},
		{
			name: "sons.of.anarchy.s05e10.480p.BluRay.x264-GAnGSteR",
			want: Torrent{
				Title:       "sons of anarchy",
				Resolution:  "480p",
				Season:      5,
				Seasons:     []int{5},
				Episode:     10,
				Source:      "bluray",
				Codec:       "x264",
				Group:       "GAnGSteR",
				ContentType: TV,
			},
		},
		{
			name: "Color.Of.Night.Unrated.DC.VostFR.BRrip.x264",
			want: Torrent{
				Season:      -1,
				Title:       "Color Of Night",
				Unrated:     true,
				Languages:   []string{"vostfr"},
				Source:      "brrip",
				Codec:       "x264",
				ContentType: Movie,
			},
		},
		{
			name: "Da Vinci Code DVDRip",
			want: Torrent{
				Season:      -1,
				Title:       "Da Vinci Code",
				Source:      "dvdrip",
				ContentType: Movie,
			},
		},
		{
			name: "Some.girls.1998.DVDRip",
			want: Torrent{
				Season:      -1,
				Title:       "Some girls",
				Source:      "dvdrip",
				Year:        1998,
				ContentType: Movie,
			},
		},
		{
			name: "Ecrit.Dans.Le.Ciel.1954.MULTI.DVDRIP.x264.AC3-gismo65",
			want: Torrent{
				Season:      -1,
				Title:       "Ecrit Dans Le Ciel",
				Source:      "dvdrip",
				Year:        1954,
				Languages:   []string{"multi"},
				Codec:       "x264",
				Audio:       "ac3",
				Group:       "gismo65",
				ContentType: Movie,
			},
		},
		{
			name: "2019 After The Fall Of New York 1983 REMASTERED BDRip x264-GHOULS",
			want: Torrent{
				Season:      -1,
				Title:       "2019 After The Fall Of New York",
				Source:      "bdrip",
				Remastered:  true,
				Year:        1983,
				Codec:       "x264",
				Group:       "GHOULS",
				ContentType: Movie,
			},
		},
		{
			name: "Ghost In The Shell 2017 720p HC HDRip X264 AC3-EVO",
			want: Torrent{
				Season:      -1,
				Title:       "Ghost In The Shell",
				Source:      "hdrip",
				Hardcoded:   true,
				Year:        2017,
				Resolution:  "720p",
				Codec:       "x264",
				Audio:       "ac3",
				Group:       "EVO",
				ContentType: Movie,
			},
		},
		{
			name: "Rogue One 2016 1080p BluRay x264-SPARKS",
			want: Torrent{
				Season:      -1,
				Title:       "Rogue One",
				Source:      "bluray",
				Year:        2016,
				Resolution:  "1080p",
				Codec:       "x264",
				Group:       "SPARKS",
				ContentType: Movie,
			},
		},
		{
			name: "Desperation 2006 Multi Pal DvdR9-TBW1973",
			want: Torrent{
				Season:      -1,
				Title:       "Desperation",
				Source:      "dvd",
				Year:        2006,
				Languages:   []string{"multi"},
				Region:      "R9",
				Group:       "TBW1973",
				ContentType: Movie,
			},
		},
		{
			name: "Maman, j'ai raté l'avion 1990 VFI 1080p BluRay DTS x265-HTG",
			want: Torrent{
				Season:      -1,
				Title:       "Maman, j'ai raté l'avion",
				Source:      "bluray",
				Year:        1990,
				Audio:       "dts",
				Resolution:  "1080p",
				Languages:   []string{"vfi"},
				Codec:       "x265",
				Group:       "HTG",
				ContentType: Movie,
			},
		},
		{
			name: "[BlurayDesuYo] Shingeki no Kyojin (Season 3) 38 (BD 1920x1080 10bit FLAC) [619BE7E0].mkv",
			want: Torrent{
				Title:       "Shingeki no Kyojin",
				Audio:       "flac",
				Resolution:  "1080",
				Container:   "mkv",
				Group:       "BlurayDesuYo",
				Season:      3,
				Seasons:     []int{3},
				Episode:     38,
				ColorDepth:  "10-bit",
				ContentType: TV,
			},
		},
		{
			name: "[Ohys-Raws] JoJo no Kimyou na Bouken Ougon no Kaze - 33 (BS11 1280x720 x264 AAC).mp4",
			want: Torrent{
				Season:      -1,
				Audio:       "aac",
				Codec:       "x264",
				Container:   "mp4",
				Episode:     33,
				Group:       "Ohys-Raws",
				Resolution:  "720",
				Title:       "JoJo no Kimyou na Bouken Ougon no Kaze",
				ContentType: TV,
			},
		},
		{
			name: "[HorribleSubs] Boruto - Naruto Next Generations - 111 [720p].mkv",
			want: Torrent{
				Season:      -1,
				Container:   "mkv",
				Episode:     111,
				Group:       "HorribleSubs",
				Resolution:  "720p",
				Title:       "Boruto - Naruto Next Generations",
				ContentType: TV,
			},
		},
		{
			name: "Marvels.Agents.of.S.H.I.E.L.D.S06E05.720p.HDTV.x264-AVS.mkv",
			want: Torrent{
				Codec:       "x264",
				Container:   "mkv",
				Episode:     5,
				Group:       "AVS",
				Resolution:  "720p",
				Season:      6,
				Seasons:     []int{6},
				Source:      "hdtv",
				Title:       "Marvels Agents of S H I E L D",
				ContentType: TV,
			},
		},
		{
			name: "stephen.colbert.2019.02.03.conan.obrien.web.x264-cookiemonster.mkv",
			want: Torrent{
				Season:      -1,
				Codec:       "x264",
				Container:   "mkv",
				Group:       "cookiemonster",
				Title:       "stephen colbert",
				Date:        "2019-02-03",
				Source:      "web",
				ContentType: TV,
			},
		},
		{
			name: "Star.Wars.Episode.IX.The.Rise.of.Skywalker.2019.2160p.WEB-DL.DDP5.1.Atmos.HEVC-BLUTONiUM.mkv",
			want: Torrent{
				Season:      -1,
				Audio:       "DDP5.1 Atmos",
				Codec:       "hevc",
				Year:        2019,
				Container:   "mkv",
				Group:       "BLUTONiUM",
				Title:       "Star Wars Episode IX The Rise of Skywalker",
				Resolution:  "4k",
				Source:      "web-dl",
				ContentType: Movie,
			},
		},
		{
			name: "Star.Wars.Episode.7.The.Force.Awakens.2015.1080p.BluRay.DTS.x264.D-Z0N3",
			want: Torrent{
				Season:      -1,
				Audio:       "dts",
				Group:       "D-Z0N3",
				Resolution:  "1080p",
				Source:      "bluray",
				Codec:       "x264",
				Title:       "Star Wars Episode 7 The Force Awakens",
				ContentType: Movie,
				Year:        2015,
			},
		},
		{
			name: "Last.Week.Tonight.with.John.Oliver.S08E01.February.14.2021.720p.HMAX.WEB-DL.DD2.0.H.264-null.mkv",
			want: Torrent{
				Audio:       "DD2.0",
				Codec:       "h264",
				Container:   "mkv",
				Group:       "null",
				Resolution:  "720p",
				Episode:     1,
				Season:      8,
				Seasons:     []int{8},
				Source:      "web-dl",
				Title:       "Last Week Tonight with John Oliver",
				ContentType: TV,
				Year:        2021,
			},
		},
		{
			name: "Pirates.of.the.Caribbean.Dead.Mans.Chest.2006.2160p.DSNP.WEB-DL.DTS-HD.MA.5.1.HDR.HEVC-WATCHER.mkv",
			want: Torrent{
				Title:       "Pirates of the Caribbean Dead Mans Chest",
				ContentType: Movie,
				Year:        2006,
				Resolution:  "4k",
				Container:   "mkv",
				Source:      "web-dl",
				Codec:       "hevc",
				Audio:       "dts-hd",
				Group:       "WATCHER",
				Season:      -1,
				Hdr:         true,
			},
		},
		{
			name: "[SubsPlease] Boku no Hero Academia - 106 (720p) [F8EFA646].mkv",
			want: Torrent{
				Title:       "Boku no Hero Academia",
				ContentType: TV,
				Resolution:  "720p",
				Container:   "mkv",
				Group:       "SubsPlease",
				Season:      -1,
				Episode:     106,
			},
		},
		{
			name: "Altered Carbon S01 2160p HDR Netflix WEBRip DD+ Atmos 5.1 x265-TrollUHD",
			want: Torrent{
				Title:       "Altered Carbon",
				ContentType: TV,
				Resolution:  "4k",
				Group:       "TrollUHD",
				Source:      "webrip",
				Codec:       "x265",
				Audio:       "DD+ Atmos",
				Season:      1,
				Seasons:     []int{1},
				Hdr:         true,
			},
		},
		{
			name: "The.Wizard.of.Oz.1939.4K.HDR.DV.2160p.BDRemux Ita Eng x265-NAHOM",
			want: Torrent{
				Title:       "The Wizard of Oz",
				ContentType: Movie,
				Year:        1939,
				Resolution:  "4k",
				Codec:       "x265",
				Group:       "NAHOM",
				Source:      "bdremux",
				Languages:   []string{"ita", "eng"},
				HdrTypes:    []string{"DV"},
				Hdr:         true,
				Season:      -1,
			},
		},
		{
			name: "Succession.S01.1080p.BluRay.REMUX.AVC.DTS-HD.MA.5.1-NOGRP",
			want: Torrent{
				Title:       "Succession",
				ContentType: TV,
				Resolution:  "1080p",
				Codec:       "avc",
				Group:       "NOGRP",
				Season:      1,
				Source:      "bluray",
				Audio:       "dts-hd",
				Seasons:     []int{1},
			},
		},
		{
			name: "Sicario 2015 PROPER 1080p BluRay DD-EX x264-TayTO.mp4",
			want: Torrent{
				Audio:       "DD-EX",
				Title:       "Sicario",
				ContentType: Movie,
				Year:        2015,
				Resolution:  "1080p",
				Codec:       "x264",
				Group:       "TayTO",
				Source:      "bluray",
				Season:      -1,
				Proper:      true,
				Container:   "mp4",
			},
		},
		{
			name: "Ant-Man.and.the.Wasp.Quantumania.2023.2160p.MA.WEB-DL.DDP5.1.Atmos.DV.HDR10.H.265-CMRG.mkv",
			want: Torrent{
				Title:       "Ant-Man and the Wasp Quantumania",
				Year:        2023,
				Resolution:  "4k",
				Container:   "mkv",
				Source:      "web-dl",
				Codec:       "h265",
				Audio:       "DDP5.1 Atmos",
				HdrTypes:    []string{"DV", "HDR10"},
				Hdr:         true,
				Season:      -1,
				Group:       "CMRG",
				ContentType: Movie,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestContentType(t *testing.T) {
	tests := []struct {
		want string
		is   ContentType
	}{
		{
			want: "Movie",
			is:   Movie,
		},
		{
			want: "TV",
			is:   TV,
		},
		{
			want: "Unknown",
			is:   Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := ContentType(tt.is); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTorrentScanAndValue(t *testing.T) {
	var torrent Torrent

	jsonStr := `{"title":"Pirates of the Caribbean Dead Mans Chest","alternativeTitle":"","contentType":0,"year":0,"resolution":"4k","extended":false,"unrated":false,"proper":false,"repack":false,"convert":false,"hardcoded":false,"retail":false,"remastered":false,"region":"","container":"mkv","source":"web-dl","codec":"hevc","audio":"dts-hd","group":"WATCHER","season":-1,"seasons":null,"episode":0,"languages":null,"hdr":true,"hdrTypes":null,"colorDepth":"","date":""}`

	err := torrent.Scan(jsonStr)
	if err != nil {
		t.Error(err)
	}
	if torrent.Title != "Pirates of the Caribbean Dead Mans Chest" {
		t.Error("Title not parsed")
	}

	str, err := torrent.Value()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, jsonStr, str)

	err = torrent.Scan([]byte(`{"title":"Rogue One A Star Wars Story","content_type":"movie","year":2016,"resolution":"1080p","container":"mkv","source":"bluray","codec":"x264","audio":"dts","group":"D-Z0N3","season":-1}`))
	if err != nil {
		t.Error(err)
	}
	if torrent.Title != "Rogue One A Star Wars Story" {
		t.Error("Title not parsed")
	}

	err = torrent.Scan(1000)
	if err == nil {
		t.Error("Expected error")
	}

}

func TestDebugParser(t *testing.T) {
	DebugParser("Star.Wars.Episode.IX.The.Rise.of.Skywalker.2019.2160p.WEB-DL.DDP5.1.Atmos.HEVC-BLUTONiUM.mkv")
}
