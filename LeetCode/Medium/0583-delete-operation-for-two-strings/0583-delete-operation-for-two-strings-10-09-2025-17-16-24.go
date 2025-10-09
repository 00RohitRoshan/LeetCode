// Memmoization
func minDistance(word1 string, word2 string) int {
    lenW1 := len(word1)
    lenW2 := len(word2)
    mem := make([][]int,lenW1)
    for i := range lenW1 {
        mem[i] = make([]int,lenW2)
        for j := range lenW2 {
            mem[i][j] = -1
        }
    }
    lcs := longestCommonSubsequence(len(word1)-1,len(word2)-1,&word1,&word2,&mem)
    return len(word1)-lcs+len(word2)-lcs
}
func longestCommonSubsequence(w1idx int, w2idx int,word1 *string, word2 *string, mem *[][]int) int{
    if w1idx < 0 || w2idx < 0 {
        return 0
    }
    if (*mem)[w1idx][w2idx] > -1 {
        return (*mem)[w1idx][w2idx]
    }
    if (*word1)[w1idx] == (*word2)[w2idx] {
        (*mem)[w1idx][w2idx] = 1+longestCommonSubsequence(w1idx-1,w2idx-1,word1,word2,mem)
        return (*mem)[w1idx][w2idx]
    }else{
        (*mem)[w1idx][w2idx] = max(
            longestCommonSubsequence(w1idx,w2idx-1,word1,word2,mem),
            longestCommonSubsequence(w1idx-1,w2idx,word1,word2,mem),
        )
        return (*mem)[w1idx][w2idx]
    }
}

// // Recurssion
// func minDistance(word1 string, word2 string) int {
//     lcs := longestCommonSubsequence(len(word1)-1,len(word2)-1,&word1,&word2)
//     return len(word1)-lcs+len(word2)-lcs
// }
// func longestCommonSubsequence(w1idx int, w2idx int,word1 *string, word2 *string) int{
//     if w1idx < 0 || w2idx < 0 {
//         return 0
//     }
//     if (*word1)[w1idx] == (*word2)[w2idx] {
//         return 1+longestCommonSubsequence(w1idx-1,w2idx-1,word1,word2)
//     }else{
//         return max(
//             longestCommonSubsequence(w1idx,w2idx-1,word1,word2),
//             longestCommonSubsequence(w1idx-1,w2idx,word1,word2),
//         )
//     }
// }