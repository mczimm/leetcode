package main

func main() {
	println(minimumLength("abaacbcbb")) //5
}

func minimumLength(s string) int {
	freq := make(map[int32]int)

	for _, v := range s {
		freq[v]++
	}

	var delCount int
	for _, v := range freq {
		if v%2 == 1 {
			delCount += v - 1
		} else {
			delCount += v - 2
		}
	}
	return len(s) - delCount
}
