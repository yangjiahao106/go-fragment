package main

func addStrings(num1 string, num2 string) string {

	ans := make([]byte, 0)
	var carry byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += num1[i] - '0'
		}
		if j >= 0 {
			carry += num2[j] - '0'
		}

		ans = append([]byte{carry%10 + '0'}, ans...)
		carry = carry / 10
	}

	return string(ans)
}
