func kthSmallest(mat [][]int, k int) int {
    sums := []int{0}
    for i := range len(mat) {
        rsums := []int{}
        for j := range len(mat[0]){
            for _, s := range sums {
                rsums = append(rsums,s+mat[i][j])
            }
        }
        sort.Ints(rsums)
        if len(rsums) > k {
            rsums = rsums[:k]
        }
        sums = rsums
    }
    return sums[k-1]
}