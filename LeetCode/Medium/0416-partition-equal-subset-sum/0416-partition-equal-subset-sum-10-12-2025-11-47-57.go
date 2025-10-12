func canPartition(nums []int) bool {
    total := 0
    for _,num := range nums {
        total += num
    }
    if total % 2 == 1 {
        return false
    }
    n := len(nums)
    mem := make([][]int,total)
    for i := range total {
        mem[i] = make([]int,n)
    }
    return partition(0,0,nums,total/2,&mem) == 2
}
func partition(idx int,sum int,nums []int,target int,mem *[][]int)int {
    if idx >= len(nums) {
        return 1
    }
    if (*mem)[sum][idx] > 0 {
        return (*mem)[sum][idx]
    }
    if sum + nums[idx] == target {
        return 2
    }
    (*mem)[sum][idx] = partition(idx+1,sum,nums,target,mem) + partition(idx+1,sum+nums[idx],nums,target,mem)
    if (*mem)[sum][idx] > 2 {
        (*mem)[sum][idx] = 2
    }
    return (*mem)[sum][idx]
}