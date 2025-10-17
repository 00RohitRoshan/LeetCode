// memoization
func maxCoins(nums []int) int {
    n := len(nums)
    dp := make([][]int,n)
    for i := range n {
        dp[i] = make([]int,n)
        for j := range n {
            dp[i][j] = -1
        }
    }
    return f(0,len(nums)-1,nums,&dp)
}
func f(i int,j int,nums []int,mem *[][]int)int{
    if i > j {
        return 0
    }
    if (*mem)[i][j] != -1  {
        return (*mem)[i][j]
    }
    front,back  := 1,1
    if i > 0 {
        front = nums[i-1]
    }
    if j < len(nums)-1 {
        back = nums[j+1]
    }
    for a := i; a<=j; a++{
        coin := front*nums[a]*back + f(i,a-1,nums,mem) + f(a+1,j,nums,mem)
        (*mem)[i][j]  = max((*mem)[i][j],coin)
    }
    return (*mem)[i][j]
}

// // recursion TLE
// func maxCoins(nums []int) int {
//     return f(0,len(nums)-1,nums)
// }
// func f(i int,j int,nums []int)int{
//     coins,front,back  := 0,1,1
//     if i > 0 {
//         front = nums[i-1]
//     }
//     if j < len(nums)-1 {
//         back = nums[j+1]
//     }
//     for a := i; a<=j; a++{
//         coin := front*nums[a]*back + f(i,a-1,nums) + f(a+1,j,nums)
//         coins  = max(coins,coin)
//     }
//     return coins
// }

// // too much mem allocation in array generation
// func maxCoins(nums []int) int {
//     n := len(nums)
//     coins := 0
//     fmt.Println(nums)
//     for i := range n {
//         l,r := 1 , 1
//         if i-1 > -1 {
//             l = nums[i-1]
//         }
//         if i+1 < n {
//             r = nums[i+1]
//         }
//         fmt.Println(l*nums[i]*r)
//         coins = max(coins,l*nums[i]*r+maxCoins(append(nums[:i],nums[i+1:]...)))
//     }
//     return coins
// }