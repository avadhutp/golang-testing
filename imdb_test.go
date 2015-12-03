package main

import (
	"github.com/eefret/gomdb"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestImdbGetMovie(t *testing.T) {
	oldGetMovie := getMovie
	defer func() { getMovie = oldGetMovie }()

	expected := new(gomdb.MovieResult)
	expected.Title = "My funny valentine"

	getMovie = func(id string) (*gomdb.MovieResult, error) {
		return expected, nil
	}

	actual, err := getMovie("some-fake-id")

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}
