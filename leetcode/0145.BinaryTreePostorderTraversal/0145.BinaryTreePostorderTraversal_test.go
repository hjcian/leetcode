package binarytreepostordertraversal

import (
	. "codepractices/leetcode/helper"
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"ex1",
			[]int{1, NULL, 2, 3},
			[]int{3, 2, 1},
		},
		{
			"ex2",
			[]int{},
			[]int{},
		},
		{
			"ex3",
			[]int{1},
			[]int{1},
		},
		{
			"ex4",
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"ex5",
			[]int{1, NULL, 2},
			[]int{2, 1},
		},
		{
			"README sample",
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		root := BuildTree(tt.nums)

		t.Run("recursive_"+tt.name, func(t *testing.T) {
			got := make([]int, 0)
			recursive(root, &got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}
		})

		t.Run("iterative_"+tt.name, func(t *testing.T) {
			if got := iterative(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
