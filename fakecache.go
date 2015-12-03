package main

var (
	determineValue = (*FakeCache).determineValue
)

// FakeCache Simple fake cache for demonstrating DI of internal struct methods
type FakeCache struct {
	value string
}

// Get get the cache value
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
