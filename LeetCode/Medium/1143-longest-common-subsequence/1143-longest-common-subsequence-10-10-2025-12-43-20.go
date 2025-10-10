// Memmoization
func longestCommonSubsequence(text1 string, text2 string) int {
    row := len(text1)
    column := len(text2)
    mem := make([][]int,row)
    for i := range row {
        mem[i] = make([]int,column)
        for j := range column {
            mem[i][j] = -1
        }
    }
    return lcs(row-1, column-1, text1, text2, &mem)
}

func lcs(idx1 int, idx2 int,text1 string, text2 string, mem *[][]int) int{
    if idx1 < 0 || idx2 < 0 {
        return 0
    }
    if (*mem)[idx1][idx2] > -1 {
        return (*mem)[idx1][idx2]
    }
    if text1[idx1] == text2[idx2] {
        (*mem)[idx1][idx2] = 1+lcs(idx1-1, idx2-1, text1, text2, mem)
        return (*mem)[idx1][idx2]
    }else{
        (*mem)[idx1][idx2] = max(
            lcs(idx1-1, idx2, text1, text2, mem),
            lcs(idx1, idx2-1, text1, text2, mem),
        )
        return (*mem)[idx1][idx2] 
    }
}

// // Recursion
// func longestCommonSubsequence(text1 string, text2 string) int {
//     return lcs(len(text1)-1, len(text2)-1, text1, text2)
// }

// func lcs(idx1 int, idx2 int,text1 string, text2 string) int{
//     if idx1 < 0 || idx2 < 0 {
//         return 0
//     }
//     if text1[idx1] == text2[idx2] {
//         return 1+lcs(idx1-1, idx2-1, text1, text2)
//     }else{
//         return max(
//             lcs(idx1-1, idx2, text1, text2),
//             lcs(idx1, idx2-1, text1, text2),
//         )
//     }
// }