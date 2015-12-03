package main

import "fmt"

func main() {
	// httpget
	g := NewGetter()
	fmt.Println(g.getData(apiURL))

	// jsondecode
	raw := "{\"param\":\"value\"}"
	fmt.Println(decodeJSON(raw))

	// imdb
	title := getTitle("tt0119116")
	fmt.Println(title)
}
