package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Helper function to get template path consistently
func getTemplatePath(filename string) string {
	templatesDir := os.Getenv("TEMPLATES_DIR")
	if templatesDir == "" {
		templatesDir = "../frontend/templates/"
	}
	return filepath.Join(templatesDir, filename)
}

// Homepage handler
func Homehandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, 404, "Page not found")
		return
	}

	artists := make([]*ArtistDetail, 0, len(appData.ArtistMap))
	for _, artist := range appData.ArtistMap {
		artists = append(artists, artist)
	}

	tmplPath := getTemplatePath("index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
		return
	}

	if err := tmpl.Execute(w, artists); err != nil {
		RenderError(w, 500, "Internal Server Error")
	}
}

// Artist page handler
func Artisthandle(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RenderError(w, 400, "Invalid artist ID")
		return
	}

	artist, exists := appData.ArtistMap[id]
	if !exists {
		RenderError(w, 404, "Artist not found")
		return
	}

	tmplPath := getTemplatePath("artist.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		RenderError(w, 500, "Internal Server Error")
	}
}

// Search handler
func Searcher(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderError(w, 405, "Method not allowed")
		return
	}

	query := strings.ToLower(strings.TrimSpace(r.FormValue("Search")))
	location := strings.ToLower(strings.TrimSpace(r.FormValue("location")))
	creationFrom, _ := strconv.Atoi(r.FormValue("creationFrom"))
	creationTo, _ := strconv.Atoi(r.FormValue("creationTo"))
	albumFrom, _ := strconv.Atoi(r.FormValue("albumFrom"))
	albumTo, _ := strconv.Atoi(r.FormValue("albumTo"))
	members, _ := strconv.Atoi(r.FormValue("members"))

	results := []*ArtistDetail{}

	for _, artist := range appData.ArtistMap {
		// SEARCH FILTER
		if query != "" {
			match := strings.Contains(strings.ToLower(artist.Name), query)
			if !match {
				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), query) {
						match = true
						break
					}
				}
			}
			if !match {
				continue
			}
		}

		// CREATION DATE FILTER
		if creationFrom != 0 && artist.CreationDate < creationFrom {
			continue
		}
		if creationTo != 0 && artist.CreationDate > creationTo {
			continue
		}

		// FIRST ALBUM FILTER
		if albumFrom != 0 || albumTo != 0 {
			parts := strings.Split(artist.FirstAlbum, "-")
			albumYear, _ := strconv.Atoi(parts[len(parts)-1])
			if albumFrom != 0 && albumYear < albumFrom {
				continue
			}
			if albumTo != 0 && albumYear > albumTo {
				continue
			}
		}

		// MEMBERS FILTER
		if members != 0 && len(artist.Members) != members {
			continue
		}

		// LOCATION FILTER
		if location != "" {
			found := false
			searchClean := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(location, ",", " "), "-", " "), "_", " ")
			for _, loc := range artist.Locations {
				locClean := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(loc), "-", " "), "_", " "), ",", " ")
				if strings.Contains(locClean, searchClean) {
					found = true
					break
				}
				searchWords := strings.Fields(searchClean)
				allWordsMatch := true
				for _, word := range searchWords {
					if !strings.Contains(locClean, word) {
						allWordsMatch = false
						break
					}
				}
				if allWordsMatch {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		results = append(results, artist)
	}

	tmplPath := getTemplatePath("index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
		return
	}

	if err := tmpl.Execute(w, results); err != nil {
		RenderError(w, 500, "Internal Server Error")
	}
}

// Error page struct
type ErrorPage struct {
	Code    int
	Message string
}

// Render error page safely
func RenderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	data := ErrorPage{
		Code:    code,
		Message: message,
	}

	tmplPath := getTemplatePath("error.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		// fallback plain text
		http.Error(w, message, code)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, message, code)
	}
}
