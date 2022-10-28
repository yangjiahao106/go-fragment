package main

func main() {
	nums := []int{1, 2, 2, 5, 5, 5, 5, 6, 7, 9, 11, 28}
	target := 3
	binarySearchWithDup(nums, target)
}

// 找到 第一个大于等于的位置
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)/2 + l
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// 二分查找包含重复值 找到第一个目标值
func binarySearchWithDup(nums []int, target int) int {
	ans := - 1
	l, r := 0, len(nums)
	for l <= r {
		m := (r-l)/2 + l
		if nums[m] == target {
			ans = m
			r = m - 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}

// 找到最后一个小于等于目标值的元素
func binarySearchLast(nums []int, target int) int {

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)>>2 + l
		if nums[mid] <= target {
			if mid < len(nums)-1 || nums[mid+1] > target { // or
				return mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return 0
}
