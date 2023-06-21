package main

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
// 1,2,3,4,5,6,7 ->  4,3,2,1,7,6,5 ->  5,6,7,1,2,3,4
// 7,6,5,4,3,2,1
func rotate(nums []int, k int) {
	k = k % len(nums)
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}

	for l, r := 0, k-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}

	for l, r := k, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}

}
