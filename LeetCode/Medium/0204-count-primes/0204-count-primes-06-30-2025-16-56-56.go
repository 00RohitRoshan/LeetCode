func countPrimes(n int) int {
    nums := make([]bool,n)
    for i := 2 ; i<n ; i++{
        nums[i] = true
    }
    for i := 2 ; i<n ; i++{
        if nums[i] {
            for j := i*i; j<n; j+=i{
                nums[j] = false
                fmt.Println(j)
            }
        }
    }
    ans := 0
    fmt.Println(nums)
    for _,B :=  range nums{
        if B {
            ans++
        }
    }
    return ans
}