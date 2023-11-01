package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28)) //AB
}

// From the soluttions
func convertToTitle(columnNumber int) string {
	if columnNumber <= 26 {
		return string(rune('A' - 1 + columnNumber))
	} else {
		return convertToTitle((columnNumber-1)/26) + string('A'+(columnNumber-1)%26)
	}
}
