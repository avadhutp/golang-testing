package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestEx1GetData(t *testing.T) {
	expected := "{\"param\":\"value\"}"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, expected)
	}))
	defer ts.Close()

	sut := NewGetter()
	actual := sut.getData(ts.URL)

	assert.Equal(t, expected, actual)
}

func TestEx1GetDataHandleError(t *testing.T) {
	sut := NewGetter()
	actual := sut.getData(httptest.DefaultRemoteAddr)

	assert.Empty(t, actual, "Invalid URL, which means http client should throw an error; therefore, we expect a nil from getData")
}
