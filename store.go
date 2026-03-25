package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

var baseURL = "https://groupietrackers.herokuapp.com/api"

func fetchArtists() ([]Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %v", err)
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("failed to decode artists: %v", err)
	}

	return artists, nil
}

func fetchLocations() ([]Location, error) {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response LocationsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}

func fetchDates() ([]Date, error) {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response DatesResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}

func fetchRelations() ([]Relation, error) {
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response RelationsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}


func loadAllData() (*AppData, error) {
	data := &AppData{
		ArtistMap: make(map[int]*ArtistDetail),
	}

	var err error

	data.Artists, err = fetchArtists()
	if err != nil {
		return nil, err
	}

	data.Locations, err = fetchLocations()
	if err != nil {
		return nil, err
	}

	data.Dates, err = fetchDates()
	if err != nil {
		return nil, err
	}

	data.Relations, err = fetchRelations()
	if err != nil {
		return nil, err
	}

	data.linkData()

	return data, nil
}



func (d *AppData) linkData() {

	locationMap := make(map[int][]string)
	for _, loc := range d.Locations {
		locationMap[loc.ID] = loc.Locations
	}

	dateMap := make(map[int][]string)
	for _, date := range d.Dates {
		dateMap[date.ID] = date.Dates
	}

	relationMap := make(map[int]map[string][]string)
	for _, rel := range d.Relations {
		relationMap[rel.ID] = rel.DatesLocations
	}

	for _, artist := range d.Artists {
		complete := &ArtistDetail{
			Artist:       artist,
			Locations:    locationMap[artist.ID],
			ConcertDates: dateMap[artist.ID],
			Relations:    relationMap[artist.ID],
		}

		d.ArtistMap[artist.ID] = complete
	}
}
