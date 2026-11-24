package main

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	cases := []struct {
		in   []int
		want [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
		{[]int{}, [][]int{{}}},
		{[]int{2, 2, 2}, [][]int{{}, {2}, {2, 2}, {2, 2, 2}}},
	}
	for _, c := range cases {
		if got := SubsetsWithDup(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("SubsetsWithDup(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
