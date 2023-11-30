package main

import "fmt"

func main() {
	// ["Shogun"]
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	// ["Shogun"]
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"}))
	// ["sad","happy"]
	fmt.Println(findRestaurant([]string{"happy", "sad", "good"},
		[]string{"sad", "happy", "good"}))
}

// From the solutions
func findRestaurant(list1 []string, list2 []string) []string {
	var res = make([]string, 1)
	var minSum = 3000

	for i, iv := range list1 {
		for j, jv := range list2 {
			if iv == jv {
				if i+j < minSum {
					minSum = i + j
					res[0] = iv
				} else if i+j == minSum {
					res = append(res, iv)
				}
			}
		}

	}
	return res
}

// From the solutions
func findRestaurant2(list1 []string, list2 []string) []string {
	restaurants1 := make(map[string]int, len(list1))
	minSum := 3000 // такого числа не может быть по условию
	res := make(map[int][]string)
	for pos, val := range list1 {
		restaurants1[val] = pos
	}
	for pos, val := range list2 {
		if posr, ok := restaurants1[val]; ok {
			if posr+pos <= minSum {
				minSum = posr + pos
				res[minSum] = append(res[minSum], val)
			}
		}
	}
	return res[minSum]
}
