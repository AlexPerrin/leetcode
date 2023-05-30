//LeetCode 1. Two Sum
//Naive Approuch using nested loops
//O(n^2) time complexity

package leetcode

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output := []int{i, j}
				return output
			}
		}
	}
	return nil
}
