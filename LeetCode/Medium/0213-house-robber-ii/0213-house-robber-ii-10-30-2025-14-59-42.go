func rob(nums []int) int {
    if len(nums) ==  1 {
        return nums[0]
    }
    return max(rec(0,nums[:len(nums)-1]),rec(0,nums[1:]))
}
func rec(i int,nums []int) int {
    if i >= len(nums) {
        return 0
    }
    return max(nums[i]+rec(i+2,nums),rec(i+1,nums))
}