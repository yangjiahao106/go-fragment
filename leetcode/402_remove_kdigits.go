package main

// 单调栈
func removeKdigits(num string, k int) string {

	stack := make([]byte, 0)
	for _, b := range []byte(num) {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > b {
			stack = stack[:len(stack)-1]
			k -= 1
		}
		stack = append(stack, b)
	}

	if k >= len(stack) {
		return "0"
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}
	return string(stack)
}
