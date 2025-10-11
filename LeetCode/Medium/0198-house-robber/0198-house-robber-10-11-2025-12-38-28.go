// memmoized
func rob(nums []int) int {
    n := len(nums)
    mem := make([]int,n)
    for i := range n {
        mem[i] = -1
    }
    return maximize(0,nums,&mem)
}

func maximize(idx int,nums []int,mem *[]int)  int {
    if idx >= len(nums) {
        return 0
    }
    if (*mem)[idx] != -1 {
        return (*mem)[idx]
    }
    (*mem)[idx] = max(
        nums[idx] + maximize(idx+2,nums,mem),
        maximize(idx+1,nums,mem),
    )
    return (*mem)[idx] 
}


// // tle recurrsion
// func rob(nums []int) int {
//     return maximize(0,nums)
// }

// func maximize(idx int,nums []int)  int {
//     if idx >= len(nums) {
//         return 0
//     }
//     return max(
//         nums[idx] + maximize(idx+2,nums),
//         maximize(idx+1,nums),
//     )
// }