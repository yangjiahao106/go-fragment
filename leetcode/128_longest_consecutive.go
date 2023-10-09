package main

import "fmt"

func TestLongestConsecutive() {

	nums := []int{100, 4, 200, 1, 3, 2} // 1,2,3,4

	ans := longestConsecutive(nums)
	fmt.Println(ans)

}

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列, 时间复杂度为 O(n)
func longestConsecutive(nums []int) int {
	val2idx := make(map[int]int)
	for i, v := range nums {
		val2idx[v] = i
	}

	marked := make([]bool, len(nums))
	ans := 0
	for i, v := range nums {
		if marked[i] {
			continue
		}

		count := 1
		tmp := v
		for {
			j, ok := val2idx[tmp-1]
			if !ok {
				break
			}
			count += 1
			marked[j] = true
			tmp--
		}

		tmp = v
		for {
			j, ok := val2idx[tmp+1]
			if !ok {
				break
			}
			count += 1
			marked[j] = true
			tmp++
		}

		if count > ans {
			ans = count
		}
	}

	return ans
}

func longestConsecutive2(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for v := range m {
		if m[v-1] {
			continue
		}
		c := v + 1
		for m[c] {
			c++
		}
		if c-v > ans {
			ans = c - v
		}
	}
	return ans
}
