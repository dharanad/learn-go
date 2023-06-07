package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	if spySleeper.calls != 3 {
		t.Errorf("not enough calls to sleeps, want 3 got %d", spySleeper.calls)
	}
}
