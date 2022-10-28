package main

import "fmt"

func TestTrap(){

	height := []int{2,8,6,4,2,0,1,5,1,5}

	fmt.Println(trap(height))
	fmt.Println(trap2(height))

}



func trap(height []int) int {

	lMax, rMax := height[0], height[len(height)-1]
	ans := 0
	l, r := 0, len(height)-1
	for l <= r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])

		if height[l] <= height[r] {
			ans += min(lMax, rMax) - height[l]
			l++
		} else {
			ans += min(lMax, rMax) - height[r]
			r--
		}

	}

	return ans
}

func trap2(height []int) int {
	stack := make([]int, 0)
	ans := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]
			w := i - left - 1
			h := min(height[left], height[i]) - top
			ans += w * h
		}
		stack = append(stack, i)

	}

	return ans
}
