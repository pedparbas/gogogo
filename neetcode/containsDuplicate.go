package neetcode

import "slices"

/*
Given an integer array *nums*, return *true* if any value appears more than once in the array, otherwise return false
*/

// O(n)
func hasDuplicate(nums []int) bool {
	arr := map[int]int{}
	for _, v := range nums {
		if arr[v] == 1 {
			return true
		}
		arr[v] = 1
	}
	return false
}

// O(n^2)
func hasDuplicateBForce(nums []int) bool {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == v {
				return true
			}
		}
	}
	return false
}

// O(nlog(n))
func hasDuplicateSorting(nums []int) bool {
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
