package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})) //2
}

func numUniqueEmails(emails []string) int {
	res := make(map[string]bool)

	for _, v := range emails {
		parts := strings.Split(v, "@")
		local := strings.Split(parts[0], "+")
		check := strings.Replace(local[0], ".", "", -1) + "@" + parts[1]

		res[check] = true
	}

	return len(res)
}
