package main

func isValid(s string) bool {

	stack := make([]byte, 0)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, b := range []byte(s) {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[b] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true

}
