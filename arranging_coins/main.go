package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(5)) //2
	fmt.Println(arrangeCoins(8)) //3
	fmt.Println(arrangeCoins(2)) //1
	fmt.Println(arrangeCoins(3)) //2
	fmt.Println()
	fmt.Println(arrangeCoins2(5)) //2
	fmt.Println(arrangeCoins2(8)) //3
	fmt.Println(arrangeCoins2(2)) //1
	fmt.Println(arrangeCoins2(3)) //2
	fmt.Println()
	fmt.Println(arrangeCoins3(5)) //2
	fmt.Println(arrangeCoins3(8)) //3
	fmt.Println(arrangeCoins3(2)) //1
	fmt.Println(arrangeCoins3(3)) //2
}

// First attempt 08.11.23
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	var res int
	tmp := n

	for i := 1; i < n; i++ {
		tmp -= i

		if tmp < 0 {
			res = i - 1
			break
		}
		if tmp == 0 {
			res = i
			break
		}
		if tmp == 1 {
			res = i
			break
		}
	}
	return res
}

// Second attempt 08.11.23
func arrangeCoins2(n int) int {
	if n == 1 {
		return 1
	}
	var res int
	tmp := n
Loop:
	for i := 1; i < n; i++ {
		tmp -= i
		switch {
		case tmp < 0:
			res = i - 1
			break Loop
		case tmp == 0:
			res = i
			break Loop
		case tmp == 1:
			res = i
			break Loop
		}
	}
	return res
}

// Third attempt 08.11.23
func arrangeCoins3(n int) int {
	if n == 1 {
		return 1
	}
	var res int
	tmp := n

	for i := 1; i < n; i++ {
		tmp -= i

		if tmp < 0 {
			res = i - 1
			break
		} else {
			res = i
		}
	}
	return res
}
