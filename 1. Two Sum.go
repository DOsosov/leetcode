package main

func twoSum(nums []int, target int) []int {
    sum := 0
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            sum = nums[i] + nums[j]
            if sum == target {
                nums[0] = i
                nums[1] = j
                
                return(nums[:2])
            }
        }
    }

    return(nums)
}