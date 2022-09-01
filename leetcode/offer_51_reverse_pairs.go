package main

import "fmt"

func TestReversePairs() {
	ans := reversePairs([]int{7, 5, 3, 2, 1, 2, 5})
	fmt.Println(ans)
}

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	count := 0
	idx := len(nums) / 2
	count += reversePairs(nums[:idx])
	count += reversePairs(nums[idx:])

	l, r := 0, idx
	temp := make([]int, 0)
	for l < idx && r < len(nums) {
		if nums[l] > nums[r] {
			temp = append(temp, nums[l])
			l++
		} else {
			temp = append(temp, nums[r])
			r++
			count += l
		}
	}
	if l < idx {
		temp = append(temp, nums[l:idx]...)
	}
	if r < len(nums) {
		temp = append(temp, nums[r:]...)
		count += (len(nums) - r) * l
	}
	for k, v := range temp {
		nums[k] = v
	}

	return count
}

func reversePairs2(nums []int) int {
	count := 0
	var merge func(nums1, nums2 []int) []int
	merge = func(nums1, nums2 []int) []int {
		res := make([]int, 0)
		i, j := 0, 0
		for i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				res = append(res, nums1[i])
				i++
			} else {
				res = append(res, nums2[j])
				j++
				count += i
			}
		}
		if i < len(nums1) {
			res = append(res, nums1[i:]...)
		}
		if j < len(nums2) {
			count += i * len(nums2[j:])
			res = append(res, nums2[j:]...)
		}
		return res
	}

	var helper func(nums []int) []int
	helper = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		m := len(nums) / 2
		left := helper(nums[:m])
		right := helper(nums[m:])
		return merge(left, right)
	}

	nums = helper(nums)
	return count
}
