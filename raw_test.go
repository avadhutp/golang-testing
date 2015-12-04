package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	True := false
	if true != True {
		t.Error("Confused Yoda is")
	}
}

func TestAssertPkg(t *testing.T) {
	expected := true
	actual := false

	assert.Equal(t, expected, actual, "This is very bad news!")
}
