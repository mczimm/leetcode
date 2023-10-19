package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))  //"5F3Z-2E9W"
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))     //"2-5G-3J"
	fmt.Println(licenseKeyFormatting2("5F3Z-2e-9-w", 4)) //"5F3Z-2E9W"
	fmt.Println(licenseKeyFormatting2("2-5g-3-J", 2))    //"2-5G-3J"
}

// My first attempt time limit exceeded
func licenseKeyFormatting(s string, k int) string {
	var res []string
	var tmp string

	//s = "0" + s[:]

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if len(tmp) < k {
			tmp = strings.ToUpper(string(s[i])) + tmp
		} else {
			res = append([]string{tmp}, res...)
			tmp = strings.ToUpper(string(s[i]))
		}
	}
	res = append([]string{tmp}, res...)
	return strings.Join(res, "-")
}

// My second attempt
func licenseKeyFormatting2(s string, k int) string {
	var tmp []byte
	var res string
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			tmp = append(tmp, s[i])
			count++
		}
		if count == k {
			tmp = append(tmp, '-')
			count = 0
		}
	}
	lenTMP := len(tmp) - 1
	//reverse
	for i := lenTMP; i >= 0; i-- {
		if i == lenTMP && tmp[i] == '-' {
			continue
		}
		res += strings.ToUpper(string(tmp[i]))
	}
	return res
}
