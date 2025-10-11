func removeDuplicates(nums []int) int {
    i,j := 0,0
    n := len(nums)
    for i<=j && j<n {
        if nums[i] != nums[j] {
            i++
            nums[i] = nums[j]
        }
        j++
    }
    return i+1
}