package main

import "fmt"

func TestGenerateMatrix() {
	ans := generateMatrix(0)
	for k := range ans {
		fmt.Println(ans[k])
	}
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	step := 1
	start, end := 0, n-1
	for start < end {
		for i := start; i < end; i++ {
			matrix[start][i] = step
			step++
		}

		for i := start; i < end; i++ {
			matrix[i][end] = step
			step++
		}

		for i := end; i > start; i-- {
			matrix[end][i] = step
			step++
		}

		for i := end; i > start; i-- {
			matrix[i][start] = step
			step++
		}
		start++
		end--
	}

	if start == end {
		matrix[start][end] = step
	}
	return matrix
}
