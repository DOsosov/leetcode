package main

func removeElement(nums []int, val int) int {
    k := 0
    for i := 0; i <= len(nums) - 1; i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k = k + 1
        }
    }
    return k
}