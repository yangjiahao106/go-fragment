package main

import "fmt"

func TestFirstMissingPositive() {
	ans := firstMissingPositive([]int{1, 1})
	fmt.Println(ans)

}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i]-1 != i {
		// nums[i] ==  nums[nums[i]-1] 时会发生死循环 例如： [1, 1],  [3,2,3]

		for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}

	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func firstMissingPositive2(nums []int) int {
	for i, n := range nums {
		for n > 0 && n <= len(nums) && i+1 != n {
			i = n - 1
			n, nums[n-1] = nums[n-1], n
		}
	}
	fmt.Println(nums)

	for i, n := range nums {
		if i+1 != n {
			return i + 1
		}
	}

	return len(nums) + 1
}
