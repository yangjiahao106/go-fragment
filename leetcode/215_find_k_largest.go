package main

import "fmt"

func testFindKthLargest(){
	for i := 1; i <= 10; i++ {
		fmt.Println(findKthLargest([]int{3,6,1,2,9,0,4,5,7,8}, i))
	}
}

func findKthLargest(nums []int, k int) int {
	var helper func(left, right int) int
	helper = func(left, right int) int {
		prov := nums[right]

		l, r := left, right
		for l < r {
			for l < r && nums[l] >= prov {
				l++
			}
			nums[r] = nums[l]
			for l < r && nums[r] < prov {
				r--
			}
			nums[l] = nums[r]
		}
		nums[l] = prov// 记得填坑

		if l == k-1 {
			return nums[l]
		}

		if l < k {
			return helper(l+1, right)
		}
		return helper(left, r-1)

	}

	return helper(0, len(nums)-1)

}

/*

*/