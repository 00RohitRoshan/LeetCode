func canPartition(nums []int) bool {
    total := 0
    for _,num := range nums {
        total += num
    }
    if total % 2 == 1 {
        return false
    }
    n := len(nums)
    mem := make([][]int,(total/2)+1)
    for i := range (total/2)+1 {
        mem[i] = make([]int,n)
    }
    return partition(0,nums,total/2,&mem) == 2
}
func partition(idx int,nums []int,target int,mem *[][]int)int {
    // fmt.Println(idx+1,nums,target,mem)
    if idx >= len(nums) || target < 0 {
        return 1
    }
    if (*mem)[target][idx] > 0 {
        return (*mem)[target][idx]
    }
    if 0 == target {
        return 2
    }
    (*mem)[target][idx] = max(partition(idx+1,nums,target,mem),partition(idx+1,nums,target-nums[idx],mem))
    return (*mem)[target][idx]
}