package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(highFive([][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}}))                                                                                                                                                                                                                                                                                                             //[[1,87],[2,88]]
	fmt.Println(highFive([][]int{{1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100}}))                                                                                                                                                                                                                                                                                                              //[[1,100],[7,100]]
	fmt.Println(highFive([][]int{{1, 46}, {1, 93}, {1, 94}, {1, 36}, {1, 71}, {2, 4}, {2, 68}, {2, 63}, {2, 80}, {2, 27}, {3, 79}, {3, 35}, {3, 95}, {3, 56}, {3, 35}, {4, 31}, {4, 32}, {4, 82}, {4, 42}, {4, 97}}))                                                                                                                                                                                                                               //[[1,68],[2,48],[3,60],[4,56]]
	fmt.Println(highFive([][]int{{1, 23}, {1, 49}, {1, 63}, {1, 66}, {1, 15}, {2, 18}, {2, 4}, {2, 32}, {2, 12}, {2, 55}, {3, 92}, {3, 78}, {3, 35}, {3, 9}, {3, 75}, {4, 4}, {4, 72}, {4, 83}, {4, 89}, {4, 48}, {5, 15}, {5, 53}, {5, 64}, {5, 85}, {5, 21}, {6, 20}, {6, 2}, {6, 20}, {6, 19}, {6, 63}, {7, 67}, {7, 40}, {7, 49}, {7, 37}, {7, 65}, {8, 86}, {8, 25}, {8, 17}, {8, 67}, {8, 16}, {9, 76}, {9, 63}, {9, 84}, {9, 93}, {9, 16}})) //[[1,43],[2,24],[3,57],[4,59],[5,47],[6,24],[7,51],[8,42],[9,66]]
}

func highFive(items [][]int) [][]int {
	students := make(map[int][]int)
	var result [][]int
	for _, item := range items {
		students[item[0]] = append(students[item[0]], item[1])
	}
	for id, sc := range students {
		score, err := topStudents(sc)
		if err != nil {
			continue
		}
		result = append(result, []int{id, score})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func topStudents(scores []int) (int, error) {
	if len(scores) == 0 {
		return 0, fmt.Errorf("err")
	}
	var avg int
	sort.Ints(scores)
	for i := len(scores) - 1; i > len(scores)-6; i-- {
		avg += scores[i]
	}
	return avg / 5, nil
}
