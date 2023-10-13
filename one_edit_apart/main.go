package main

import "fmt"

func main() {
	fmt.Println(oneEditApart("cat", "dog"))   //false
	fmt.Println(oneEditApart("cat", "cats"))  //true
	fmt.Println(oneEditApart("cat", "cut"))   //true
	fmt.Println(oneEditApart("c", "cut"))     //false
	fmt.Println(oneEditApart("task", "cast")) //false
	fmt.Println(oneEditApart("cat", "at"))    //true
	fmt.Println(oneEditApart("cat", "a"))     //false
	fmt.Println(oneEditApart("cat", "acts"))  //false
	fmt.Println(oneEditApart("scat", "acts")) //false
	fmt.Println(oneEditApart("cat", "act"))   //false
}

func oneEditApart(first string, second string) bool {
	if len(first) < len(second) {
		first, second = second, first
	}

	if len(first)-len(second) > 1 { // check if the one string not bigger than 2 symbols another string
		return false
	}

	detect := false
	if len(first) == len(second) { // check one by one symbols if the string are the same length
		for i := range first {
			if first[i] != second[i] {
				if detect {
					return false
				}
				detect = true
			}
		}
	} else { // loop over by shortest string and check current and next symbol of long string
		for i := 0; i < len(second); i++ {
			if detect {
				if first[i+1] != second[i] {
					return false
				}
			} else {
				if first[i] != second[i] {
					detect = true
				}
			}
		}
	}
	return true
}
