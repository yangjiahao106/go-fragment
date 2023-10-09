package main

import "fmt"

func TestLengthOfLIS() {

	fmt.Println(lengthOfLISTemp([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISTemp([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLISTemp([]int{0, 1, 0, 3, 2, 3, 9, 1, 5, 8}))

}

func lengthOfLISTemp(nums []int) int {
	dp := make([]int, 0)

	for _, v := range nums {
		if len(dp) == 0 || v > dp[len(dp)-1] {
			dp = append(dp, v)
			continue
		}

		l, r := 0, len(dp)-1
		for l < r { // 找到第一个 大于等于n的位置，并替换
			m := (l + r) / 2
			if dp[m] >= v {
				r = m
			} else {
				l = m + 1
			}
		}
		dp[r] = v

	}
	return len(dp)
}

// dp 60ms
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

// 贪心 + 二分 8ms
func lengthOfLIS2(nums []int) int {
	d := make([]int, 0)
	var tmp []int
	_ = tmp

	for _, n := range nums {
		if len(d) > 0 && n <= d[len(d)-1] {
			l, r := 0, len(d)
			for l < r { // 找到第一个 大于等于n的位置，并替换
				mid := (l + r) / 2
				if d[mid] >= n {
					r = mid
				} else {
					l = mid + 1
				}
			}
			d[l] = n
		} else {
			d = append(d, n)
			tmp = d
		}
	}

	fmt.Println(tmp)
	return len(d)
}

/*
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        d = []
        for n in nums:
            if not d or n > d[-1]:
                d.append(n)
            else:
                l, r = 0, len(d) - 1
                loc = r
                while l <= r:
                    mid = (l + r) // 2
                    if d[mid] >= n:
                        loc = mid
                        r = mid - 1
                    else:
                        l = mid + 1
                d[loc] = n
        return len(d)

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-by-leetcode-soluti/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
