package binary_subarrays_with_sum

import "testing"

func TestSubarraysWithSum(t *testing.T) {
	result := numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestSubarraysWithSumHandlesEmptyArray(t *testing.T) {
	result := numSubarraysWithSum([]int{}, 2)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestSubarraysWithSumHandlesAllOnes(t *testing.T) {
	result := numSubarraysWithSum([]int{1, 1, 1, 1, 1}, 3)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestSubarraysWithSumHandlesAllZeros(t *testing.T) {
	result := numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0)
	if result != 15 {
		t.Errorf("Expected 15, but got %d", result)
	}
}
