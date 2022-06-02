package main

func twoSum(nums []int, target int) []int {
	var numsLength int = len(nums)
	for i := 0; i < numsLength; i++ {
		for j := 0; j < numsLength; j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
