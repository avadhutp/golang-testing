package main

import "fmt"

func main() {
	// httpget
	g := NewGetter()
	fmt.Println(g.getData(apiURL))

	// jsondecode
	raw := "{\"param\":\"value\"}"
	fmt.Println(decodeJSON(raw))
}
