package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // HL
)

func TestAssertPkg(t *testing.T) {
	expected := true
	actual := false

	assert.Equal(t, expected, actual, "This is very bad news!") // HL
}
