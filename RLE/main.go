package main

import "fmt"

/*
	два указателя
	цикл по строке
	если ук1 == ук2, то увелич стетчик
	иначе если счетчие больше 1, то добавить в результат букву и цифру
	иначе тольео букву
*/

func main() {
	str := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBBXX"
	var res string
	right := 1
	cnt := 1
	var last_ch string

	for left := 0; left < len(str)-1; left++ {
		if str[left] == str[right] {
			cnt++
		} else {
			if cnt > 1 {
				res += string(str[left]) + fmt.Sprint(cnt)
				cnt = 1
			} else {
				res += string(str[left])
			}
		}
		right++
	}

	last_ch = string(str[len(str)-1:])

	if cnt > 1 {
		res += last_ch + fmt.Sprint(cnt)
	} else {
		res += last_ch
	}
	fmt.Println(res) //A4B3C2XYZD4E3F3A6B28

}
