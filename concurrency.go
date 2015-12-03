package main

var process = (*FakeRoutine).process

type Metric struct {
	name  string
	value float64
}

type FakeRoutine struct {
}

func (f *FakeRoutine) sendMetric(m *Metric) {
	if m.value != 0 {
		go process(f, m)
	}
}

func (f *FakeRoutine) process(m *Metric) {
	// Does some processing
}
