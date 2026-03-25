package main

import (
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {

	data := &AppData{
		ArtistMap: make(map[int]*ArtistDetail),
	}

	Routes(data)

	req, _ := http.NewRequest("GET", "/", nil)

	if req == nil {
		t.Error("request should not be nil")
	}
}