package main

import "fmt"

func testDailyTemperatures() {
	ans := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(ans)

}

func dailyTemperatures(temperatures []int) []int {

	stack := make([]int, 0)
	ans := make([]int, len(temperatures))
	for k, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			i := stack[len(stack)-1]
			ans[i] = k - i

			stack = stack[:len(stack)-1]

		}
		stack = append(stack, k)
	}
	return ans
}
