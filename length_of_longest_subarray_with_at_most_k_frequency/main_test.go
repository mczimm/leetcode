package length_of_longest_subarray_with_at_most_k_frequency

import "testing"

func TestMaxSubarrayLengthWith6Frequencies(t *testing.T) {
	result := maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestMaxSubarrayLengthWith2Frequency(t *testing.T) {
	result := maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestMaxSubarrayLengthWith4SameElements(t *testing.T) {
	result := maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestMaxSubarrayLengthWith3DifferentElements(t *testing.T) {
	result := maxSubarrayLength([]int{1, 1, 1, 3}, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestMaxSubarrayLengthWithEmptyArray(t *testing.T) {
	result := maxSubarrayLength([]int{}, 2)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestMaxSubarrayLengthWithZeroFrequency(t *testing.T) {
	result := maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 0)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}
