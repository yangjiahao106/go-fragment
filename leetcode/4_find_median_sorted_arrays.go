package main

import "fmt"

func TestFindMedianSortedArrays() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}

// log(m + n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 查找第k小元素
	var helper func(k int) int
	helper = func(k int) int {
		idx1, idx2 := 0, 0
		for {
			if idx1 == len(nums1) {
				return nums2[idx2+k-1]
			}
			if idx2 == len(nums2) {
				return nums1[idx1+k-1]
			}
			if k == 1 {
				return min(nums1[idx1], nums2[idx2])
			}

			newIdx1 := min(idx1+k/2-1, len(nums1)-1)
			newIdx2 := min(idx2+k/2-1, len(nums2)-1)

			if nums1[newIdx1] <= nums2[newIdx2] {
				newIdx1 += 1
				k -= newIdx1 - idx1
				idx1 = newIdx1
			} else {
				newIdx2 += 1
				k -= newIdx2 - idx2
				idx2 = newIdx2
			}
		}
	}

	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return float64(helper(totalLen/2 + 1))
	}
	return (float64(helper(totalLen/2)) + float64(helper(totalLen/2+1))) / 2

}
