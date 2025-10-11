func maxSumAfterPartitioning(arr []int, k int) int {
    n := len(arr)
    mem := make([][]int,n)
    for i := range n {
        mem[i] = make([]int,n)
        for j := range n {
            mem[i][j] = -1
        }
    }
    return Partition(0,0,0,arr,k,&mem)
}
func Partition(startIdx int,endIdx int,maxEle int,arr []int,k int,mem *[][]int)  int {
    if endIdx >= len(arr) {
        return 0
    } 
    if (*mem)[startIdx][endIdx] != -1 {
        return (*mem)[startIdx][endIdx]
    }
    maxEle = max(maxEle,arr[endIdx])
    if endIdx-startIdx+1 == k {
        (*mem)[startIdx][endIdx] = maxEle*k+Partition(endIdx+1,endIdx+1,0,arr,k,mem)
        return (*mem)[startIdx][endIdx]
    }
    (*mem)[startIdx][endIdx] = max(
        maxEle*(endIdx-startIdx+1)+Partition(endIdx+1,endIdx+1,0,arr,k,mem),
        Partition(startIdx,endIdx+1,maxEle,arr,k,mem),
    )
    return (*mem)[startIdx][endIdx]
}


// // this is wrong as I thought I can change the order
// func maxSumAfterPartitioning(arr []int, k int) int {
//     sort.Ints(arr)
//     n := len(arr)
//     numpartition := math.Ceil(float64(n)/float64(k))
//     arr = arr[n-int(numpartition):]
//     sum := 0
//     for i:= len(arr)-1; i>= 0; i-- {
//         if n > k {
//             n -= k
//             sum += k*arr[i]
//         }else{
//             sum += n*arr[i]
//         }
//     }
//     return sum
// }