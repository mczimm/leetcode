package main

import "fmt"

func main() {
	//"Sao Paulo"
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	//"A"
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
}

func destCity(paths [][]string) string {
	var res string
	depart := make(map[string]struct{})
	dest := make(map[string]struct{})

	for i := 0; i < len(paths); i++ {
		depart[paths[i][0]] = struct{}{}
		dest[paths[i][1]] = struct{}{}
	}

	for i := range dest {
		if _, ok := depart[i]; !ok {
			res = i
		}
	}

	return res
}
