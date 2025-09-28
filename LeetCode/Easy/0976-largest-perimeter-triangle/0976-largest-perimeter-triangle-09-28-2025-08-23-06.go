func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    len := len(nums)
    for n := len; n > 2; n--{
        a := nums[n-3]
        b := nums[n-2]
        c := nums[n-1]
        if a+b > c {
            return a+b+c
        }
    }
    return 0
}