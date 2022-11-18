package main

func maxSlidingWindow(nums []int, k int) []int {

	q := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	ans := make([]int, 0)
	ans = append(ans, nums[q[0]])

	for i := k; i < len(nums); i++ {
		if nums[i-k] == nums[q[0]] {
			q = q[1:]
		}
		// 两种方式都可以
		//if q[0] <= i-k {
		//	q = q[1:]
		//}

		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		ans = append(ans, nums[q[0]])
	}

	return ans
}
