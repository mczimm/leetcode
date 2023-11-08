package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(addStrings("11", "123"))                                  //134
	fmt.Println(addStrings("3876620623801494171", "6529364523802684779")) //10405985147604178950
	fmt.Println()
	fmt.Println(addStrings2("11", "123"))                                  //134
	fmt.Println(addStrings2("3876620623801494171", "6529364523802684779")) //10405985147604178950
}

// First attempt
func addStrings(num1 string, num2 string) string {
	n1, _ := strconv.ParseInt(num1, 10, 64)
	n2, _ := strconv.ParseInt(num2, 10, 64)
	n3 := n1 + n2
	fmt.Println(n3)
	return ""
}

// Second attempt
func addStrings2(num1 string, num2 string) string {
	bigIntNum1, _ := new(big.Int).SetString(num1, 10)
	bigIntNum2, _ := new(big.Int).SetString(num2, 10)

	sum := new(big.Int) //.Add(bigIntNum1, bigIntNum2)
	sum.Add(bigIntNum1, bigIntNum2)

	return sum.String()
}
