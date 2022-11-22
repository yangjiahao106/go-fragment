package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func TestSubGroup() {

	GroupList := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}
	ans, err := Subgroup(GroupList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)

}

func Subgroup(groups [][]string) ([][]string, error) {
	// 检查参数
	if len(groups) <= 1 {
		return nil, errors.New("groups not valid")
	}
	for i := 1; i < len(groups); i++ {
		if len(groups[i]) != len(groups[0]) {
			return nil, errors.New("groups not valid")
		}
	}

	// 打乱所有成员顺序保证结果的随机性
	shuffle(groups)
	for i := 0; i < len(groups); i++ {
		shuffle2(groups[i])
	}

	ans := make([][]string, 0)

	l := 0
	length := len(groups)

	for {
		sub := make([]string, 0, 2)
		i := l
		for ; i < l+length; i++ {
			if len(groups[i%length]) == 0 {
				continue
			}
			sub = append(sub, groups[i%length][0])
			groups[i%length] = groups[i%length][1:]
			if len(sub) >= 2 {
				ans = append(ans, sub)
				break
			}
		}

		l = (i + 1) % length
		if len(sub) == 0 {
			break
		}

		if len(sub) == 1 && len(ans) > 0 {
			ans[len(ans)-1] = append(ans[len(ans)-1], sub...)
		}
	}

	return ans, nil
}

// Shuffle 打乱顺序保证随机
func shuffle(slice [][]string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func shuffle2(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
