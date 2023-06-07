package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dharan")
	got := buffer.String()
	want := "Hello, Dharan"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
