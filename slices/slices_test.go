package slices

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("sum of all elements in a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d , wanted %d, given %v", got, want, numbers)
		}
	})

	t.Run("sum of all given slices with elements in them", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum of all tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		// tail = sum of all elements in a slice except the first one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty tail sum", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 3, 9})
		want := []int{0, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
