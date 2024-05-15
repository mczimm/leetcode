package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	// Existing test case
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	result := searchInsert(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Additional test cases
	nums = []int{1, 3, 5, 6}
	target = 2
	expected = 1
	result = searchInsert(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	nums = []int{1, 3, 5, 6}
	target = 7
	expected = 4
	result = searchInsert(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	nums = []int{1, 3, 5, 6}
	target = 0
	expected = 0
	result = searchInsert(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
