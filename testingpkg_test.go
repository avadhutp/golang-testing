package main

import "testing" // HL

func TestMe(t *testing.T) {
	True := false
	if true != True {
		t.Error("Confused Yoda is")
	}
}
