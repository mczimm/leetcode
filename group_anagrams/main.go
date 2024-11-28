package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//[["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//[["max"],["buy"],["doc"],["may"],["ill"],["duh"],["tin"],["bar"],["pew"],["cab"]]
	fmt.Println(groupAnagrams2([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
	fmt.Println(groupAnagrams2([]string{"duh", "ill"})) //same sum
}

// My first attempt
func groupAnagrams(strs []string) [][]string {
	tp := make(map[[26]int][]string)
	var res [][]string

	for i := 0; i < len(strs); i++ {
		var ts [26]int
		// I've given it from solution
		for j := 0; j < len(strs[i]); j++ {
			ts[strs[i][j]-'a']++
		}
		// End
		tp[ts] = append(tp[ts], strs[i])
	}
	for _, v := range tp {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	tmp := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		tmp[key] = append(tmp[key], str)
	}
	var res [][]string
	for _, i := range tmp {
		res = append(res, i)
	}
	return res
}

func sortString(in string) string {
	s := strings.Split(in, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
