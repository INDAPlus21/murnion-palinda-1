package sum

import "testing"

func TestSumEvenArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected_result := 55

	result := ConcurrentSum(array, 5)

	if result != expected_result {
		t.Errorf("could not sum even array")
	}
}

func TestSumOddArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected_result := 45

	result := ConcurrentSum(array, 5)

	if result != expected_result {
		t.Errorf("could not sum even array")
	}
}

func TestSumNegatives(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -4, -5, -6}
	expected_result := 30

	result := ConcurrentSum(array, 5)

	if result != expected_result {
		t.Errorf("could not sum even array")
	}
}