package main

import "github.com/eefret/gomdb"

var (
	getMovie = gomdb.MovieByImdbID
)

func getTitle(id string) string {
	if r, err := getMovie(id); err == nil {
		return r.Title
	}

	return ""
}
