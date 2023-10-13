package main

import "fmt"

func main() {
	//fmt.Println(checkRecord("PPALLP")) //true
	//fmt.Println(checkRecord("PPALLL")) //false
	fmt.Println(checkRecord("AA"))   //false
	fmt.Println(checkRecord("LALL")) //true

}

func checkRecord(s string) bool {
	a := 0
	l := 0
	for _, v := range s {
		if a > 1 {
			return false
		}
		if l >= 3 {
			return false
		}
		if string(v) == "A" {
			a++
			l = 0
		} else if string(v) == "L" {
			l++
		} else {
			l = 0
		}
	}
	if l >= 3 || a > 1 {
		return false
	}
	return true
}
