package main

import (
	"fmt"
	"sort"
)

func TestSearchMatrix() {
	/*
		每行中的整数从左到右按升序排列。
		每行的第一个整数大于前一行的最后一个整数。 (拉成一维依然有序)
	*/

	nums := []int{1, 3, 5, 8, 29}
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 0
	})
	fmt.Println(i)


	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	ans := searchMatrix2(matrix, 1)
	fmt.Println(ans)
}

func searchMatrix(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	// 拉成一维

	row := len(matrix)
	col := len(matrix[0])

	l, r := 0, row*col-1
	for l <= r {
		//fmt.Println(l, r)
		mid := (l + r) / 2
		v := matrix[mid/col][mid%col]

		if v == target {
			return true
		}
		if v > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func searchMatrix3(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)

	return col < len(matrix[row]) && matrix[row][col] == target
}

func searchMatrix4(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
	return i < m*n && matrix[i/n][i%n] == target
}
