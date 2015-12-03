package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

// TestJsonDecode Demonstrates parametrized tests
func TestJsonDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]string
		msg      string
	}{
		{"{\"phil\":\"coulson\"}", map[string]string{"phil": "coulson"}, "Valid JSON input should provide a valid map"},
		{"{\"phil\":\"coulson}", nil, "Invalid JSON input should provide nil"},
	}

	for _, test := range tests {
		actual := decodeJSON(test.input)

		assert.Equal(t, test.expected, actual, test.msg)
	}
}
