func findTargetSumWays(nums []int, target int) int {
    return compute(nums,target,len(nums)-1)
}
func compute(nums []int, target int, idx int) int {
    if idx == -1 && target == 0 {
        return 1
    }else if idx == -1 && target != 0{
        return 0
    }
    return compute(nums,target+nums[idx],idx-1) + compute(nums,target-nums[idx],idx-1)
}