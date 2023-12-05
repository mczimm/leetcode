package main

import "fmt"

func main() {
	fmt.Println(isPathCrossing("NES"))   // false
	fmt.Println(isPathCrossing("NESWW")) // true
	fmt.Println(isPathCrossing("SS"))    // false
}

// My first attempt, from the Editorial algorithm
func isPathCrossing(path string) bool {
	coordinates := make(map[byte][]int)
	visited := make(map[string]struct{})
	var x, y int

	visited["0:0"] = struct{}{}
	coordinates['N'] = []int{0, 1}
	coordinates['S'] = []int{0, -1}
	coordinates['W'] = []int{-1, 0}
	coordinates['E'] = []int{1, 0}

	for i := 0; i < len(path); i++ {
		t := coordinates[path[i]]
		x += t[0]
		y += t[1]
		if _, ok := visited[fmt.Sprintf("%d:%d", x, y)]; ok {
			return true
		}
		visited[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
	}

	return false
}
