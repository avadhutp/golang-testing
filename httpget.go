package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiURL = "http://jsonplaceholder.typicode.com/posts/1"
)

// Getter Test struct for url content retreival
type Getter struct {
	client http.Client
}

// NewGetter OOP constructor for testin
func NewGetter() *Getter {
	g := new(Getter)
	g.client = http.Client{Timeout: 10 * time.Second}

	return g
}

func (g *Getter) getData(url string) string {
	r, err := g.client.Get(url)

	if err != nil {
		return ""
	}

	defer r.Body.Close()

	if raw, err := ioutil.ReadAll(r.Body); err == nil {
		return string(raw)
	}

	return ""
}
