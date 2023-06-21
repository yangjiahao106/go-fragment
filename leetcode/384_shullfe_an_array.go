package main

import "math/rand"

type Solution struct {
	nums     []int
	shuffled []int
}

func Constructor_(nums []int) Solution {
	return Solution{
		nums:     nums,
		shuffled: append([]int{}, nums...),
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	//rand.Shuffle(len(this.nums), func(i, j int) {
	//	this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	//})
	// 相当于不放回地随机抽剩余的数字，抽取后放入最后面
	for i := len(this.shuffled) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		this.shuffled[i], this.shuffled[r] = this.shuffled[r], this.shuffled[i]
	}

	return this.shuffled
}
