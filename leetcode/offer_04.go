package main

import "fmt"

func testFindNumberIn2DArray() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	for i := 0; i < 31; i++ {
		ans := findNumberIn2DArray(matrix, i)
		fmt.Println(i, ans)

	}

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row <= len(matrix)-1 && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}

	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	colMax := len(matrix[0]) - 1
	for colMax > 0 && matrix[0][colMax] > target {
		colMax--
	}

	rowMax := len(matrix) - 1
	for rowMax > 0 && matrix[rowMax][0] > target {
		rowMax--
	}

	for col := 0; col <= colMax; col++ {
		l, r := 0, rowMax
		for l <= r {
			m := (l + r) / 2
			if matrix[m][col] == target {
				return true
			}
			if matrix[m][col] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return false
}
