package main

func main() {
	println(doesValidArrayExist([]int{1, 1, 0})) // true
}

func doesValidArrayExist(derived []int) bool {
	var sum int
	for _, v := range derived {
		sum += v
	}
	return sum%2 == 0
}
