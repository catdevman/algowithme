package twosum

import (
	"fmt"
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{0, 2},
		},
	}

	for _, test := range tests {
		out := TwoSum(test.nums, test.target)
		if len(out) != len(test.expected) {
			t.Fatal(fmt.Sprintf("output of TwoSum had len %d but expected has len %d", len(out), len(test.expected)))
		}
		for _, expected := range test.expected {
			if !slices.Contains[[]int](out, expected) {
				t.Fatal(fmt.Sprintf("output does not have %d like it should", expected))
			}
		}
	}
}
