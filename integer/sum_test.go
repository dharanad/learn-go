package integer

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 3, 4})
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		want := []int{5, 7}
		checkSums(t, got, want)
	})
	t.Run("make sum of slices with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})
}
