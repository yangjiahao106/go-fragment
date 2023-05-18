package main

import "fmt"

func TestSortColors() {
	nums := []int{2, 2, 2, 0, 1, 2, 0, 1, 0, 0, 2, 2}
	sortColors3(nums)
	fmt.Println(nums)
}

/*
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l = 0
        for i in [0,1,2]:
            for r in range(l, len(nums)):
                if nums[r] == i :
                    nums[l], nums[r] = nums[r], nums[l]
                    l += 1

*/

func sortColors(nums []int) {
	// 常数空间 + 扫描一趟
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0 += 1
		}

		if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1 += 1
		}

		if p1 < p0 {
			p1 = p0
		}
	}
}

func sortColors2(nums []int) {
	// 遍历两趟
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p += 1
		}
	}

	for i := p; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[p], nums[i] = nums[i], nums[p]
			p += 1
		}
	}
}

// 官方题解
func sortColors3(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
