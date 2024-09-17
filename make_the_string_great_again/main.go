package main

import "fmt"

func main() {
	fmt.Println(makeGood("leEeetcode")) //leeetcode
	fmt.Println(makeGood("abBAcC"))     //
	fmt.Println(makeGood("s"))          //s
	fmt.Println(makeGood("Pp"))         //
}

func makeGood(s string) string {
	var stack []rune

	for _, v := range s {
		//Thus we can tell that two characters make a pair,
		//when and only when their ASCII values differ by 32 (Since the sentence only contains letters of alphabet,
		//we do not need to consider about other speical characters). Keep This is a very common trick, keep it in mind!
		if len(stack) > 0 && (stack[len(stack)-1]-v == 32 || stack[len(stack)-1]-v == -32) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
