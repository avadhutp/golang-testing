package main

import (
	"fmt"

	"github.com/eefret/gomdb"
)

var (
	getMovie = gomdb.MovieByImdbID
)

func getTitle(id string) string {
	if r, err := getMovie(id); err == nil {
		return r.Title
	}

	return ""
}

func main() {
	fmt.Println("Input: tt0119116")
	fmt.Println("Output: " + getTitle("tt0119116"))
}
