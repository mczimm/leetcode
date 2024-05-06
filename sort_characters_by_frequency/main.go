package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	key rune
	val int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].val < p[j].val }

func main() {
	fmt.Println(frequencySort("tree")) //eert
}

func frequencySort(s string) string {
	cntM := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		cntM[rune(s[i])]++
	}
	res := sortMapByValues(cntM)

	sb := strings.Builder{}

	for _, v := range res {
		for i := 0; i < v.val; i++ {
			sb.WriteRune(v.key)
		}
	}
	return sb.String()
}

func sortMapByValues(m map[rune]int) PairList {
	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	//sort.Sort(p)
	sort.Sort(sort.Reverse(p))
	return p
}
