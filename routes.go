package main

import (
	"net/http"
)

var appData *AppData

func Routes(data *AppData) {
	appData = data

	http.HandleFunc("/", Homehandle)
	http.HandleFunc("/artist", Artisthandle)
	http.HandleFunc("/search", Searcher)

	//  Serve static files from frontend/static
	fs := http.FileServer(http.Dir("../frontend/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func StartServer(port string) error {
	return http.ListenAndServe(port, nil)
}