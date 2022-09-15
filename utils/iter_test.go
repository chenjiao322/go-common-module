package utils

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	ans := Map(func(t int) int { return t + 1 }, []int{1, 1, 1, 2, 34})
	fmt.Println(ans)
}

func TestGroupBy(t *testing.T) {
	ans := GroupBy(func(t int) int { return t }, []int{1, 1, 1, 2, 34})
	fmt.Println(ans)
}

func TestFilter(t *testing.T) {
	ans := Filter(func(t int) bool { return t > 0 }, []int{-1, 1, 1, -2, 34})
	fmt.Println(ans)
}

func TestFold(t *testing.T) {
	ans := Fold(func(t int, Acc int) int { return Acc + t }, 0, []int{-1, 1, 1, -2, 34})
	fmt.Println(ans)
}

func TestForeach(t *testing.T) {
	ForEach(func(i int) { fmt.Println(i) }, []int{1, 2, 3})
}

func TestZip(t *testing.T) {
	fmt.Println(Zip([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}))
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3}
	Reverse(a)
	fmt.Println(a)
}
