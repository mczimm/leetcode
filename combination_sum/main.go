package main

func main() {
	println(combinationSum([]int{3, 4, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int
	var combinations []int
	backtrack(target, combinations, 0, candidates, &results)
	return results
}

func backtrack(target int, combinations []int, start int, candidates []int, results *[][]int) {
	newCombinations := make([]int, len(combinations))
	copy(newCombinations, combinations)
	if target == 0 {
		*results = append(*results, newCombinations)
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		combinations = append(combinations, candidates[i])
		backtrack(target-candidates[i], combinations, i, candidates, results)
		combinations = combinations[:len(combinations)-1]
	}
}
