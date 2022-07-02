package main

import (
	"develop/algorithm/sort/bubble"
	"develop/algorithm/sort/heap"
	"develop/algorithm/sort/insert"
	"develop/algorithm/sort/merge"
	"develop/algorithm/sort/selection"
	"fmt"
)

func main() {
	// n := []int{3, 2, 1, 5, 4, 9, 12}
	n := []int{9, 8, 5, 4, 9, 1, 15, 13, 21, 3}
	b := bubble.Sort(n)
	s := selection.Sort(n)
	i := insert.Sort(n)
	h := heap.Sort(n)
	m := merge.Sort(n)

	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(i)
	fmt.Println(h)
	fmt.Println(m)
}
