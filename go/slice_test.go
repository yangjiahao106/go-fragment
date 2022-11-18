package main

import (
	"fmt"
	"testing"
)

type bigInt struct {
	i int32
	j int8
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	s = append(s, []int{1, 2, 3, 4, 5}...)
	t.Log(len(s), cap(s))
	//runtime.GOMAXPROCS()
	for i := 0; i < 4098; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestNewSlice(t *testing.T) {
	s1 := make([]int, 0)
	s2 := new([]int)

	s1 = append(s1, 1)
	*s2 = append(*s2, 1)
	fmt.Println(s1, *s2)

}
