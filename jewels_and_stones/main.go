package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) //3
}

// My first attempt
func numJewelsInStones(jewels string, stones string) int {
	jewelsD := make(map[byte]struct{})
	var res int

	for i := 0; i < len(jewels); i++ {
		jewelsD[jewels[i]] = struct{}{}
	}

	for i := 0; i < len(stones); i++ {
		if _, ok := jewelsD[stones[i]]; ok {
			res++
		}
	}
	return res

}
