package main

import "testing"

func TestLinkData(t *testing.T) {

	data := &AppData{
		Artists: []Artist{
			{
				ID: 1,
				Name: "Queen",
			},
		},
		Locations: []Location{
			{
				ID: 1,
				Locations: []string{"london-uk"},
			},
		},
		Dates: []Date{
			{
				ID: 1,
				Dates: []string{"1973-07-13"},
			},
		},
		Relations: []Relation{
			{
				ID: 1,
				DatesLocations: map[string][]string{
					"london-uk": {"1973-07-13"},
				},
			},
		},
		ArtistMap: make(map[int]*ArtistDetail),
	}

	data.linkData()

	if len(data.ArtistMap) == 0 {
		t.Error("ArtistMap should not be empty after linking data")
	}
}