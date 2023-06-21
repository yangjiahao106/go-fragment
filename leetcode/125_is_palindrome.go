package main

import "strings"

func isPalindrome_(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isNum(s[l]) {
			l++
		}
		for l < r && !isNum(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isNum(v byte) bool {
	if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {
		return true
	}
	return false

}
