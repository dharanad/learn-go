package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add small numbers", func(t *testing.T) {
		got := Add(1, 2)
		want := 3
		assertIntegers(t, got, want)
	})

	t.Run("add large numbers", func(t *testing.T) {
		got := Add(1e9, 1)
		want := int(1e9 + 1)
		assertIntegers(t, got, want)
	})

	t.Run("add negative numbers", func(t *testing.T) {
		got := Add(-11, -2)
		want := -13
		assertIntegers(t, got, want)
	})

	t.Run("add negative numbers & positive number", func(t *testing.T) {
		got := Add(-11, 2)
		want := -9
		assertIntegers(t, got, want)
	})
}

func assertIntegers(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// ExampleAdd is a example function. Examples are compiled
// (and optionally executed) as part of a package's test suite.
// Ref: https://go.dev/blog/examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
