package main

func main() {
	//println(wordBreak("leetcode", []string{"leet", "code"}))                       //true
	//println(wordBreak("applepenapple", []string{"apple", "pen"}))                  //true
	//println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) //false
	println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"})) //true
}

func wordBreak2(s string, wordDict []string) bool {
	var cur string
	//var ans string
	var i int
	//var mid string
	d := make(map[string]string)
	for _, v := range wordDict {
		d[v] = v
	}

	for _, v := range s {
		cur += string(v)
		i++
		if _, ok := d[cur]; ok {
			//mid += cur
			cur = ""
			s = s[i:]
			i = 0
		}
	}

	//for _, v := range s {
	//	cur += string(v)
	//	if _, ok := d[cur]; ok {
	//delete(d, cur)
	//mid += cur
	//cur = ""
	//}
	//}
	//if cur != "" {
	//	if _, ok := d[ans+cur]; ok {
	//		ans += cur
	//	}
	//}
	//return len(ans) == len(s)
	return len(s) == 0
}

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
