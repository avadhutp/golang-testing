package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

// START PART 1 OMIT
func TestFakeCacheGet(t *testing.T) {
	oldDetermineValue := determineValue
	defer func() { determineValue = oldDetermineValue }()

	tests := []struct {
		initialValue           string
		expected               string
		isDetermineValueCalled bool
		msg                    string
	}{
		{"", "I'm driving with the top down.", true, "Initial value is nil, so determineValue must be called"},
		{"ola", "ola", false, "Initial value is set, so determineValue must not be called"},
	}

	for _, test := range tests {
		determineValueCalled := false
		determineValue = func(f *FakeCache) {
			determineValueCalled = true
			f.value = test.expected
		}
		// END PART 1 OMIT
		// START PART 2 OMIT

		sut := new(FakeCache)
		sut.value = test.initialValue

		actual := sut.Get()

		assert.Equal(t, test.isDetermineValueCalled, determineValueCalled, test.msg)
		assert.Equal(t, test.expected, actual)
	}
}

//END PART 2 OMIT
