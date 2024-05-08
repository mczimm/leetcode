package number_of_good_pairs

import "testing"

func TestIdenticalPairsInArray1(t *testing.T) {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestIdenticalPairsInArray2(t *testing.T) {
	result := numIdenticalPairs([]int{1, 1, 1, 1})
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestIdenticalPairsInArray3(t *testing.T) {
	result := numIdenticalPairs([]int{1, 2, 3})
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}
