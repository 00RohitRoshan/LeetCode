func longestPalindrome(s string) string {
    n := len(s)
    res := ""
    for i := range n-1 {
        if s[i] == s[i+1] {
            p := expand(i,i+1,s)
            if len(p) > len(res) {
                res = p
            }
        }
        p := expand(i,i,s)
        if len(p) > len(res) {
            res = p
        }
    }
    return res
}
func expand(left int, right int, s string) string {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left+1 : right]
}

// // TC O(n^2)
// func longestPalindrome(s string) string {
//     maxs := ""
//     n := len(s)
//     for i := range n {
//         for j := i; j<=n; j++ {
//             subs := s[i:j]
//             revs := reverse(subs)
//             // fmt.Println(maxs,subs,i,j)
//             if subs == revs && len(revs) > len(maxs) {
//                 maxs = subs
//             }
//         }
//     }
//     return maxs
// }

// func reverse(s string) string {
//     r := []rune(s)
//     n := len(r)
//     for i := range n/2 {
//         r[i],r[n-i-1] = r[n-i-1],r[i]
//     }
//     return string(r)
// }