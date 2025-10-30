func subsets(nums []int) [][]int {
    n := len(nums)
    subset := 1<<n
    res := [][]int{}
    for i := range subset {
        set := []int{}
        for j := range n {
            if (i & (1<<j)) > 0 {
                set = append(set,nums[j])
            }
        }
        res = append(res,set)
    }
    return res
}