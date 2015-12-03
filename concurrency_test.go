package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestConcurrencySendMetric(t *testing.T) {
	oldProcess := process
	defer func() { process = oldProcess }()

	c := make(chan *Metric)
	defer close(c)

	processCalled := false
	process = func(f *FakeRoutine, m *Metric) {
		processCalled = true
		c <- m
	}

	expected := &Metric{"test.metric", 3.14}
	sut := new(FakeRoutine)

	sut.sendMetric(expected)

	<-c

	assert.True(t, processCalled)
}
