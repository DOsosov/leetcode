package main

func searchInsert(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == target {
			return i
		}
		if (target < nums[i]) && (target > nums[i-1]) {
			return i
		}

	}
	return 0
}