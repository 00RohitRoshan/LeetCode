func ladderLength(beginWord string, endWord string, wordList []string) int {
    has := make(map[string]bool)
    for _,i := range wordList {
        has[i] = true
    }
    if has[endWord] == false {
        return 0
    }
    q := []string{beginWord}
    l := 0
    for len(q) > 0 {
        lenQ := len(q)
        l++
        for i := range lenQ {
            wordByte := []byte(q[i])
            for i,v := range wordByte {
                for c := byte('a'); c <= byte('z'); c++{
                    if c != v {
                        wordByte[i] = c
                        newS := string(wordByte)
                        if has[newS] {
                            q = append(q,newS)
                            delete(has, newS)
                            if newS == endWord {
                                return l+1
                            }
                        }
                    }
                }
                wordByte[i] = v
            }
        }
        q = q[lenQ:]
    }
    return 0
    // return bfs([]string{beginWord},endWord,has)
}
//  func bfs(q []string, ew string, has map[string]bool) int {
//     cs := "abcdefghijklmnopqrstuvwxyz"
//     lenq := len(q)
//     for _,i := range q {
//         wBytes := []byte(i)
//         for j,v := range wBytes {
//             for k := range 25 {
//                 if v != cs[k] {
//                     wBytes[j] = cs[k]
//                     n := string(wBytes)
//                     if has[n] {
//                         q = append(q,n)
//                     }
//                     if n == ew {
//                         return 1
//                     }
//                 }
//             }
//         }
//     }
//     q = q[lenq:]
//     return bfs(q,ew,has)+1
//  }