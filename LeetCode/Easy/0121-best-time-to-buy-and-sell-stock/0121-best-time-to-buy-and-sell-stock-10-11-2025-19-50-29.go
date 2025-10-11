// tracked min buy price
func maxProfit(prices []int) int {
    n := len(prices)
    mem := make([]int,n)
    mem[0] = prices[0]
    maxProfit := 0 
    for i := 1; i<n; i++ {
        maxProfit = max(maxProfit,prices[i]-mem[i-1])
        mem[i] = min(mem[i-1],prices[i])
        // fmt.Printf("mem %+v maxprofit %d i %d\n",mem,maxProfit,i)
    }
    return maxProfit
}



// // n^2 TLE
// func maxProfit(prices []int) int {
//     maxProfit := 0
//     n := len(prices)
//     for i := range n  {
//         for j := i; j<n; j++{
//             if prices[j] - prices[i] > maxProfit {
//                 maxProfit = prices[j] - prices[i]
//             }
//         }
//     }
//     return maxProfit
// }