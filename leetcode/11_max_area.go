package main

import "fmt"

func TestMaxAres() {

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		h := min(height[l], height[r])
		w := r - l
		ans = max(ans, h*w)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// unqualified
func maxArea2(height []int) int {
	stack := make([]int, 0)
	ans := 0
	for r, v := range height {
		for _, l := range stack {
			h := min(height[l], v)
			w := r - l
			if h*w > ans {
				ans = h * w
			}
		}

		if len(stack) == 0 || v > height[stack[len(stack)-1]] {
			stack = append(stack, r)
		}
	}

	return ans
}


