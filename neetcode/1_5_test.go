package neetcode

import (
	"slices"
	"testing"
)

type Test[T, K any] struct {
	test          T
	expectedValue K
}

var (
	TEST_SUITE_MSG = "Testing %s implementation %d"
	TEST_MSG       = "Testing for %v. [Expected: %v, Received: %v]"
)

func TestContainsDuplicate(test *testing.T) {
	tests := []Test[[]int, bool]{
		{[]int{1, 2, 3, 3}, true},
		{[]int{1, 2, 3, 4}, false},
	}

	for i, f := range []func([]int) bool{hasDuplicate, hasDuplicateBForce, hasDuplicateSorting} {
		test.Logf(TEST_SUITE_MSG, "ContainsDuplicate", i+1)
		for _, t := range tests {
			res := f(t.test)
			test.Logf(TEST_MSG, t.test, t.expectedValue, res)
			if res != t.expectedValue {
				test.Fail()
			}
		}
	}
}

func TestIsAnagram(test *testing.T) {
	tests := []Test[[2]string, bool]{
		{[...]string{"racecar", "carrace"}, true},
		{[...]string{"jar", "jam"}, false},
	}

	for i, f := range []func(string, string) bool{isAnagram, isAnagramSorting} {
		test.Logf(TEST_SUITE_MSG, "IsAnagram", i+1)
		for _, t := range tests {
			res := f(t.test[0], t.test[1])
			test.Logf(TEST_MSG, t, t.expectedValue, res)
			if res != t.expectedValue {
				test.Fail()
			}
		}
	}
}

type TwoSumInput struct {
	nums   []int
	target int
}

func TestTwoSums(test *testing.T) {
	tests := []Test[TwoSumInput, []int]{
		{TwoSumInput{[]int{3, 4, 5, 6}, 7}, []int{0, 1}},
		{TwoSumInput{[]int{4, 5, 6}, 10}, []int{0, 2}},
	}

	for i, f := range []func([]int, int) []int{twoSum, twoSumsBForce} {
		test.Logf(TEST_SUITE_MSG, "TwoSums", i+1)
		for _, t := range tests {
			// var res []int
			res := f(t.test.nums, t.test.target)
			test.Logf(TEST_MSG, t, t.expectedValue, res)
			slices.Sort(res)
			for i := range res {
				if res[i] != t.expectedValue[i] {
					test.Fail()
					return
				}
			}
		}
	}

}
