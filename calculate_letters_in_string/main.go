/*
Дана строка (возможно, пустая), содержащая буквы A-Z:
AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB
Нужно написать функцию RLE, которая выведет строку вида:
A4B3C2XYZD4E3F3A6B28
Выдавать ошибку, если на вход приходит недопустимая строка.
Примечания:
Если символ встречается один раз, он остается неизменным;
Если символ встречается более одного раза, к нему добавляется число повторений.
*/

package main

import (
	"fmt"
)

func main() {
	l := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	var res []string
	right := 1
	cnt := 1

	//todo: check incorrect string

	for left := 0; left < len(l)-1; left++ {
		if l[left] == l[right] {
			cnt++
			right++
			continue
		} else if cnt > 1 {
			res = append(res, fmt.Sprintf("%s:%d", string(l[left]), cnt))
			cnt = 1
			right++
		} else {
			res = append(res, string(l[left]))
			right++
		}
	}
	if cnt > 1 {
		res = append(res, fmt.Sprintf("%s:%d", string(l[right-1]), cnt))
	} else {
		res = append(res, string(l[right-1]))
	}

	fmt.Println(res)
	fmt.Println(main2())
}

func main2() []string {
	str := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	var res []string
	right := 1
	count := 1

	//todo: check incorrect string

	for left := 0; left < len(str)-1; left++ {
		if str[left] == str[right] {
			count++
			right++
			continue
		} else {
			if count == 1 {
				res = append(res, string(str[left]))
			} else {
				res = append(res, fmt.Sprintf("%s:%d", string(str[left]), count))
			}
			count = 1
			right++
		}
	}

	if count == 1 {
		res = append(res, string(str[right-1]))
	} else {
		res = append(res, fmt.Sprintf("%s:%d", string(str[right-1]), count))
	}

	return res
}
