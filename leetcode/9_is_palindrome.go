package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 0 && x%10 == 0 { // 判断一下 0结尾的数字
		return false
	}
	tmp := 10
	for tmp < x {
		tmp = tmp*10 + x%10
		x = x / 10
	}

	return tmp == x || tmp/10 == x
}
