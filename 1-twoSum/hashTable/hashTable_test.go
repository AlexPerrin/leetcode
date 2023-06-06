package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type TwoSumTest struct {
	nums     []int
	target   int
	expected []int
}

var TwoSumTests = []TwoSumTest{
	TwoSumTest{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	TwoSumTest{[]int{3, 2, 4}, 6, []int{1, 2}},
	TwoSumTest{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range TwoSumTests {
		fmt.Println("Output", TwoSum(test.nums, test.target), "expected", test.expected)
		if !reflect.DeepEqual(TwoSum(test.nums, test.target), test.expected) { //if not equal
			t.Errorf("Output %q not equal to expected %q", TwoSum(test.nums, test.target), test.expected)
		}
	}
}
