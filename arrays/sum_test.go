package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}

	got := Sum(numbers)
	expected := 21

	if expected != got {
		t.Errorf("Expected %d got %d given %v", expected, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	firstNumbers := []int{1, 2, 3, 4, 5, 6}
	secondNumbers := []int{5}

	got := SumAll(firstNumbers, secondNumbers)
	expected := []int{21, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v got %v", expected, got)
		}
	}

	t.Run("should calculate the sum of the tails", func(t *testing.T) {
		firstNumbers := []int{1, 2, 3, 4, 5, 6}
		secondNumbers := []int{5, 5}

		got := SumAllTails(firstNumbers, secondNumbers)
		expected := []int{20, 5}

		checkSums(t, got, expected)
	})

	t.Run("should safely sum empty lists", func(t *testing.T) {
		firstNumbers := []int{1, 2, 3, 4, 5, 6}
		secondNumbers := []int{}

		got := SumAllTails(firstNumbers, secondNumbers)
		expected := []int{20, 0}

		checkSums(t, got, expected)
	})
}
