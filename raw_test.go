package main

import "testing"

func TestRaw(t *testing.T) {
	x := 1
	y := 2

	if x != y {
		t.Errorf("Expected %d; got %d", x, y)
	}
}
