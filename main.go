package main

import "fmt"

func main() {
	// Ex1: HTTP Url fetch testing
	g := NewGetter()
	fmt.Println(g.getData(apiURL))
}
