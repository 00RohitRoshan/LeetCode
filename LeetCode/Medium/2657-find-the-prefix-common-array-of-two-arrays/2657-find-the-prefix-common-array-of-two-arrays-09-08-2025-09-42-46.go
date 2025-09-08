func findThePrefixCommonArray(A []int, B []int) []int {
    inA := make(map[int]bool)
    inB := make(map[int]bool)
    ans := make([]int,len(A))
        count := 0
    for i := range len(A){
        if A[i] == B[i] {
            count += 1
        }else{
            _,okA := inA[B[i]]
            _,okB := inB[A[i]]
            if okA {
                count += 1
            }
            if okB {
                count += 1
            }
        }
        ans[i] = count
        inA[A[i]] = true
        inB[B[i]] = true
    }
    return ans
}