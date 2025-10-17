func missingNumber(nums []int) int {
    res := len(nums)
    for i := range len(nums) {
        res ^= i
        res ^= nums[i]
    }
    return res
}