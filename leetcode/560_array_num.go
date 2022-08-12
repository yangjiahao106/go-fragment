package main

import "fmt"

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	ans := 0
	for _, v := range nums {
		sum = v + sum
		if n, ok := m[sum-k]; ok {
			ans += n
		}
		m[sum] = m[sum] + 1
	}

	return ans
}

func subarraySumTest() {
	fmt.Println(subarraySum([]int{1, -1, 3, 2, 1, 3}, 3))

}
