package main

var (
	determineValue = (*FakeCache).determineValue
)

type FakeCache struct {
	value string
}

func (f *FakeCache) Get() string {
	if f.value == "" {
		determineValue(f)
	}

	return f.value
}

func (f *FakeCache) determineValue() {
	// Some computing here to calculate the actual value
	f.value = "I am iron man"
}
