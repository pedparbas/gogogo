package neetcode

/*

given an array of integers *nums* and an integer *target*, return the indices *i, j* such that *nums[i] + nums[j] == target* and *i!=j*

*/

// O(n)
func twoSum(nums []int, target int) []int {
	r := make(map[int]int)
	for i, v := range nums {
		val := target - v
		if found, ok := r[v]; ok && found != i {
			return []int{found, i}
		}
		r[val] = i
	}
	// we assume that there's always a solution
	return []int{0, 0}
}

func twoSumsBForce(nums []int, target int) []int {
	for i, v := range nums {
		for j, x := range nums {
			if j != i && v+x == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
