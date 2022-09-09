package main

import (
	"fmt"
	"math/rand"
)

func TestGetLeastNumbers() {
	for i := 0; ; i++ {
		nums := []int{10, 0, 11, 6, 3, 2, 8, 7, 9, 4, 1, 5}
		if i > len(nums) {
			return
		}
		ans := getLeastNumbers(nums, i)
		fmt.Println(ans)

	}
}

func getLeastNumbers(arr []int, k int) []int {
	var helper func(left, right int)
	helper = func(left, right int) {
		if left >= right {
			return
		}
		n := left + rand.Intn(right-left)
		arr[n], arr[right] = arr[right], arr[n]
		p := arr[right]
		l, r := left, right
		for l < r {
			for l < r && arr[l] <= p {
				l++
			}
			arr[r] = arr[l]
			for l < r && arr[r] > p {
				r--
			}
			arr[l] = arr[r]
		}

		arr[l] = p

		if l > k-1 {
			helper(left, l-1)
		} else if l < k-1 {
			helper(l+1, right)
		} else {
			return
		}
	}

	helper(0, len(arr)-1)
	return arr[:k]
}
