func wordBreak(s string, wordDict []string) bool {
    // Create a set of words for faster lookup
    wordSet := make(map[string]struct{})
    for _, word := range wordDict {
        wordSet[word] = struct{}{}
    }

    // Memoization table
    mem := make([]int, len(s)+1)
    for i := range mem {
        mem[i] = -1
    }

    // Helper function for recursion with memoization
    var canBreak func(startIdx int) int
    canBreak = func(startIdx int) int {
        // If we've reached the end, it's a valid split
        if startIdx == len(s) {
            return 1
        }

        // If we've already computed the result, return it
        if mem[startIdx] != -1 {
            return mem[startIdx]
        }

        // Try every substring starting from startIdx
        for endIdx := startIdx + 1; endIdx <= len(s); endIdx++ {
            if _, ok := wordSet[s[startIdx:endIdx]]; ok {
                // If the substring is a valid word, check the remaining string
                if canBreak(endIdx) == 1 {
                    mem[startIdx] = 1
                    return 1
                }
            }
        }

        mem[startIdx] = 0
        return 0
    }

    // Start recursion from index 0
    return canBreak(0) == 1
}


// // accepted
// func wordBreak(s string, wordDict []string) bool {
//     has := make(map[string]struct{})
//     for _,word := range wordDict {
//         has[word] = struct{}{}
//     }
//     mem := make([][]int,len(s))
//     for i := range len(s) {
//         mem[i] = make([]int,len(s)+2)
//         for j := range len(s)+2 {
//             mem[i][j] = -1
//         }
//     }
//     return canBreak(0,0,s,has,&mem) == 1
// }
// func canBreak(startIdx int,endIdx int,s string,has map[string]struct{},mem *[][]int) int {
//     if endIdx > len(s) {
//         (*mem)[startIdx][endIdx] = 0
//         return 0
//     }
//     _,ok := has[s[startIdx:endIdx]]
//     // fmt.Println(s[startIdx:endIdx])
//     if endIdx == len(s) && ok {
//         (*mem)[startIdx][endIdx] = 1
//         return 1
//     }
//     if (*mem)[startIdx][endIdx] != -1 {
//         return (*mem)[startIdx][endIdx] 
//     }
//     if ok {
//         (*mem)[startIdx][endIdx] = canBreak(endIdx,endIdx+1,s,has,mem) + canBreak(startIdx,endIdx+1,s,has,mem)
//         if (*mem)[startIdx][endIdx] > 1 {
//             (*mem)[startIdx][endIdx] = 1
//         }
//         return (*mem)[startIdx][endIdx] 
//     }
//     (*mem)[startIdx][endIdx] = canBreak(startIdx,endIdx+1,s,has,mem)
//     return (*mem)[startIdx][endIdx] 
// }


// // edge case when "aa" and "aaa" also available or "abc" and "abcd"
// func wordBreak(s string, wordDict []string) bool {
//     i := 0
//     j := 0
//     has := map[string]struct{}{}
//     for _,s := range wordDict {
//         has[s] = struct{}{}
//     }
//     fmt.Println(has)
//     for i < len(s) && j < len(s)+1 {
//         fmt.Printf("%d %d\n",i,j)
//         if _,ok := has[s[i:j]]; ok {
//             fmt.Println(s[i:j])
//             i = j
//         }
//         j++
//     }
//     return i == len(s)
// }