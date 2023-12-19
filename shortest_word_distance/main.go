package main

import (
	"fmt"
	"math"
)

func main() {
	//3
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	//1
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
	//1
	fmt.Println(shortestDistance([]string{"a", "a", "b", "b"}, "a", "b"))
	//7
	fmt.Println(shortestDistance([]string{"this", "is", "a", "long", "sentence", "is", "fun", "day", "today", "sunny", "weather", "is", "a", "day", "tuesday", "this", "sentence", "run", "running", "rainy"}, "weather", "long"))
	//1
	fmt.Println(shortestDistance([]string{"a", "b"}, "a", "b"))
	//3
	fmt.Println(shortestDistance([]string{"a", "b", "c", "d", "d"}, "a", "d"))
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	var res = float64(len(wordsDict))
	var idx1, idx2 = -1, -1

	if len(wordsDict) == 2 {
		return 1
	}

	for i := 0; i < len(wordsDict); i++ {
		if wordsDict[i] == word1 {
			idx1 = i
		} else if wordsDict[i] == word2 {
			idx2 = i
		}
		if idx1 != -1 && idx2 != -1 {
			tmp := float64(idx1 - idx2)
			res = math.Min(res, math.Abs(tmp))
		}
	}
	return int(res)
}
