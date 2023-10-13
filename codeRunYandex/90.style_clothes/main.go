package main

import (
	"fmt"
	"log"
)

func main() {
	var numtShort int
	_, err := fmt.Scan(&numtShort)
	if err != nil {
		log.Fatal(err)
	}

	var colortShort = make([]int, numtShort)
	for i := 0; i < numtShort; i++ {
		_, err := fmt.Scan(&colortShort[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	var numPants int
	_, err = fmt.Scan(&numPants)
	if err != nil {
		log.Fatal(err)
	}

	var colorPants = make([]int, numPants)
	for i := 0; i < numPants; i++ {
		_, err = fmt.Scan(&colorPants[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	for numtShort > numPants {
		colorPants = append(colorPants, 0)
		numPants++
	}

	for numtShort < numPants {
		colortShort = append(colortShort, 0)
		numtShort++
	}

	res := style(colortShort, colorPants)
	fmt.Println(res[0], res[1])
}

func style(tShortColor []int, pantsColor []int) []int {
	var res = make([]int, 2)
	i, j := 0, 0

	if len(tShortColor) == 1 && len(pantsColor) == 1 {
		return []int{tShortColor[0], pantsColor[0]}
	}

	for {
		if tShortColor[i] > pantsColor[j] {
			j++
		} else if tShortColor[i] < pantsColor[j] {
			i++
		}
		res[0] = tShortColor[i]
		res[1] = pantsColor[j]

		if res[0] == res[1] {
			return res
		}

		if i >= len(tShortColor)-1 {
			break
		}
		if j >= len(pantsColor)-1 {
			break
		}
	}

	return res
}
