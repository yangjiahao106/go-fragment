package main

func findDiagonalOrder(mat [][]int) []int {
	ans := make([]int, 0)

	flat := true
	for i := 0; i < len(mat); i++ {
		layer := getLayer(mat, i, 0)
		flat = !flat
		if flat {
			reverse(layer)
		}
		ans = append(ans, layer...)
	}

	for j := 1; j < len(mat[0]); j++ {
		layer := getLayer(mat, len(mat)-1, j)
		flat = !flat
		if flat {
			reverse(layer)
		}
		ans = append(ans, layer...)
	}
	return ans
}

func getLayer(mat [][]int, i, j int) []int {
	layer := make([]int, 0)
	for i >= 0 && j < len(mat[0]) {
		layer = append(layer, mat[i][j])
		i--
		j++
	}
	return layer
}

func reverse(layer []int) {
	for i, j := 0, len(layer)-1; i < j; i, j = i+1, j-1 {
		layer[i], layer[j] = layer[j], layer[i]
	}
}
