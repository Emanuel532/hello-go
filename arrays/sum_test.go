package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, numbers %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3, 5})
	want := []int{3, 10}
	assertEqualSlices(t, got, want)

}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum slices tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 5, 7})
		want := []int{5, 15}
		assertEqualSlices(t, got, want)
	})
	t.Run("Safely sum slices with no tail(s)", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 5, 7})
		want := []int{0, 15}
		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
