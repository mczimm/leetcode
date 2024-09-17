package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))                           // /home
	fmt.Println(simplifyPath("/home//foo/"))                      // /home/foo
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures")) // /home/user/Pictures
	fmt.Println(simplifyPath("/../"))                             // /
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))            // /.../b/d
}

func simplifyPath(path string) string {
	splited := strings.Split(path, "/")

	var stack []string

	for _, v := range splited {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
