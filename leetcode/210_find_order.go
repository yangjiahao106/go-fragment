package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {

	rely := map[int][]int{} // 课程 to 依赖它的课程
	grade := map[int]int{}  // 课程 to 入度

	for _, p := range prerequisites {
		rely[p[1]] = append(rely[p[1]], p[0])
		grade[p[0]] += 1 // 有依赖 入度加1
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if grade[i] == 0 { // 从入度为0 的开始学习
			queue = append(queue, i)
		}
	}

	ans := make([]int, 0) // 学习顺序
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		ans = append(ans, c)

		for _, v := range rely[c] {
			grade[v] -= 1
			if grade[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}

	return []int{}
}

func TestFindOrder() {

	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))

}
