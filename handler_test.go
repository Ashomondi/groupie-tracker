package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// setupTestTemplates sets TEMPLATES_DIR for testing
func setupTestTemplates() {
	cwd, _ := os.Getwd()
	templatesDir := filepath.Join(cwd, "../frontend/templates")
	os.Setenv("TEMPLATES_DIR", templatesDir)
}

// setupTestAppData initializes a minimal appData for testing
func setupTestAppData() {
	appData = &AppData{
		Artists: []Artist{
			{ID: 1, Name: "Test Artist"},
		},
		Locations: []Location{},
		Dates:     []Date{},
		Relations: []Relation{},
		ArtistMap: map[int]*ArtistDetail{
			1: &ArtistDetail{
				Artist: Artist{
					ID:   1,
					Name: "Test Artist",
				},
				Locations:    []string{},
				ConcertDates: []string{},
				Relations:    map[string][]string{},
			},
		},
	}
}

// Test the Home page handler
func TestHomehandle(t *testing.T) {
	setupTestTemplates()
	setupTestAppData()

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Homehandle)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", rr.Code)
	}
}

// Test the 404 handler when page does not exist
func TestHomehandle404(t *testing.T) {
	setupTestTemplates()
	setupTestAppData()

	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Homehandle)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 404 but got %d", rr.Code)
	}
}