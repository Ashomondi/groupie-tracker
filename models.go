package main

// holds urls
type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"` //url
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Locations    string   `json:"locations"` //url
    ConcertDates string   `json:"concertDates"` //url
    Relations    string   `json:"relations"` //url
}
// holds the actual data fetched from the json
type ArtistDetail struct {
	Artist
	Locations []string
	ConcertDates []string
	Relations map[string][]string
}
// Location represents concert locations
type Location struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
    Dates     string   `json:"dates"` 
}

// Date represents concert dates
type Date struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}

// Relation links locations and dates and artist
type Relation struct {
    ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

type AppData struct {
    Artists     []Artist
    Locations   []Location
    Dates       []Date
    Relations   []Relation
    ArtistMap   map[int]*ArtistDetail
}

type LocationsResponse struct {
    Index []Location `json:"index"`
}

type DatesResponse struct {
    Index []Date `json:"index"`
}

type RelationsResponse struct {
    Index []Relation `json:"index"`
}