package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 6, 7, 7, 8}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) []int {
	k := 1
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i-1] != nums[i] {
			nums[k] = nums[i]
			k = k + 1
		}
	}
	return nums[:k]
}
