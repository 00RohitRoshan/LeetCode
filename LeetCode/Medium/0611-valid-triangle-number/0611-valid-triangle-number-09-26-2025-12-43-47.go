func triangleNumber(nums []int) int {
    n := len(nums) 
    sort.Ints(nums)
    if n <  3 {
        return 0
    }
    count := 0
    for i := n-1; i >= 2; i-- {
        j,k := 0 , i-1
        for j < k {
            if nums[j] + nums[k] > nums[i] {
                // fmt.Printf("%d %d %d - ",i,j,k)
                count += (k-j)
                k--
            }else{
                j++
            }
        }
    }
    return count
}