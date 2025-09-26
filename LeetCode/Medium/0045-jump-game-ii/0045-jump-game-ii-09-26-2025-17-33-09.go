func jump(nums []int) int {
    n := len(nums)
    mem := make([]int,n)
    for i := range n-1 {
        mem[i] = 99999
    }
    for i := n-2; i >= 0; i-- {
        for j := 0 ; j <= nums[i]; j++ {
            c := i+j
            if c < n {
                mem[i] = min(mem[i],mem[c]+1)
            }
        }
        
    }
    return mem[0]
}
func rec(nums []int,cur int)int{
    if cur == len(nums)-1 {
        return 0
    }
    mJump := 99999
    for i := 1; i <= nums[cur]; i++  {
        mJump = min(rec(nums,cur+i),mJump)
    }
    return mJump+1
}