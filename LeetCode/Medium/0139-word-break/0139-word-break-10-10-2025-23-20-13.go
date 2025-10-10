func wordBreak(s string, wordDict []string) bool {
    has := make(map[string]struct{})
    for _,word := range wordDict {
        has[word] = struct{}{}
    }
    mem := make([][]int,len(s))
    for i := range len(s) {
        mem[i] = make([]int,len(s)+2)
        for j := range len(s)+2 {
            mem[i][j] = -1
        }
    }
    return canBreak(0,0,s,has,&mem) == 1
}
func canBreak(startIdx int,endIdx int,s string,has map[string]struct{},mem *[][]int) int {
    if endIdx > len(s) {
        (*mem)[startIdx][endIdx] = 0
        return 0
    }
    _,ok := has[s[startIdx:endIdx]]
    // fmt.Println(s[startIdx:endIdx])
    if endIdx == len(s) && ok {
        (*mem)[startIdx][endIdx] = 1
        return 1
    }
    if (*mem)[startIdx][endIdx] != -1 {
        return (*mem)[startIdx][endIdx] 
    }
    if ok {
        (*mem)[startIdx][endIdx] = canBreak(endIdx,endIdx+1,s,has,mem) + canBreak(startIdx,endIdx+1,s,has,mem)
        if (*mem)[startIdx][endIdx] > 1 {
            (*mem)[startIdx][endIdx] = 1
        }
        return (*mem)[startIdx][endIdx] 
    }
    (*mem)[startIdx][endIdx] = canBreak(startIdx,endIdx+1,s,has,mem)
    return (*mem)[startIdx][endIdx] 
}


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