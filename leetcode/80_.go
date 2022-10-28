package main

import "fmt"

func TestRemoveDuplicates() {
	nums := []int{1, 1, 1, 1, 2, 3, 3, 3, 4, 5, 6, 6, 6}
	ans := removeDuplicates(nums)
	fmt.Println(nums[:ans], ans)

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	s, f := 2, 2
	for f < len(nums) {
		if nums[f] != nums[s-2] {
			nums[s] = nums[f]
			s++
			f++
		} else {
			f++
		}
	}

	return s
}


