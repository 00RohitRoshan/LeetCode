func arrayNesting(nums []int) int {
    idNum := make(map[int]int)
    for i , j := range nums {
        idNum[i]=j
    }
    maxLen := math.MinInt
    has := make(map[int]bool)
    // sort.Ints(nums)
    for i := range len(nums) {
        maxLen = max(maxLen,build(has,i,nums,0))
    }
    return maxLen
}
func build(has map[int]bool,i int,nums []int,count int)int{
    if _,ok := has[nums[i]] ; ok {
        return count
    }
    if has[nums[i]] {
        return 0
    }
    has[nums[i]] = true
    count++
    n := build(has,nums[i],nums,count)
    return n
}