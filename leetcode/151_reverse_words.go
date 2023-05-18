package main

import "fmt"

func reverseWordsTest() {
	ans := reverseWords("hello world fuck")
	fmt.Println(ans)
}

func reverseWords(s string) string {
	b := []byte(s)
	// 反转每个单词
	for i := 0; i < len(b); {
		if b[i] == ' ' {
			i++
			continue
		}
		j := i
		for j+1 < len(b) && b[j+1] != ' ' {
			j++
		}
		temp := j
		for ; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		i = temp + 1
	}

	// 反转
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	// 去除多余空格
	slow := 0
	for f := 0; f < len(b); f++ {
		if b[f] == ' ' && (f == 0 || b[f-1] == ' ') {
			continue
		}
		b[slow] = b[f]
		slow++
	}
	// 去除尾部多余空格
	if slow > 0 && b[slow-1] == ' ' {
		slow--
	}

	return string(b[:slow])
}
