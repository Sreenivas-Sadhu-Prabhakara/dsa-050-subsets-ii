package main

import (
	"fmt"
	"sort"
)

// SubsetsWithDup returns all unique subsets of nums (which may contain
// duplicates), each subset sorted, subsets in lexicographic order.
// Sorting first lets one skip rule suppress duplicate subsets at the source.
func SubsetsWithDup(nums []int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	out := [][]int{}
	cur := []int{}
	var walk func(start int)
	walk = func(start int) {
		out = append(out, append([]int{}, cur...))
		for i := start; i < len(a); i++ {
			if i > start && a[i] == a[i-1] {
				continue // same value at the same depth would rebuild the same subset
			}
			cur = append(cur, a[i])
			walk(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	walk(0)
	return out
}

func main() {
	fmt.Println(SubsetsWithDup([]int{1, 2, 2}))
}
