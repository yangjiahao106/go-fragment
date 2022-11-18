package main

import (
	"fmt"
	"strings"
)

func RestoreIpAddressTest() {
	ans := restoreIpAddresses("25525511135")
	fmt.Println(ans)
}

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	var dfs func(left string, path []string)
	dfs = func(left string, path []string) {
		if left == "" {
			if len(path) == 4 {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}

		v := 0
		for i := 0; i < 3; i++ {
			if i >= len(left) {
				return
			}
			v = v*10 + int(left[i]-'0')
			if v <= 255 {
				dfs(left[i+1:], append(path, left[:i+1]))
			}
			if v == 0 {
				return
			}
		}

	}

	dfs(s, []string{})

	return ans
}
