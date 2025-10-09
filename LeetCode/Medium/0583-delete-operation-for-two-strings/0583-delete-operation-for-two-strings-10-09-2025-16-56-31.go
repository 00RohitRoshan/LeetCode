func minDistance(word1 string, word2 string) int {
    lcs := longestCommonSubsequence(len(word1)-1,len(word2)-1,&word1,&word2)
    return len(word1)-lcs+len(word2)-lcs
}
func longestCommonSubsequence(w1idx int, w2idx int,word1 *string, word2 *string) int{
    if w1idx < 0 || w2idx < 0 {
        return 0
    }
    if (*word1)[w1idx] == (*word2)[w2idx] {
        return 1+longestCommonSubsequence(w1idx-1,w2idx-1,word1,word2)
    }else{
        return max(
            longestCommonSubsequence(w1idx,w2idx-1,word1,word2),
            longestCommonSubsequence(w1idx-1,w2idx,word1,word2),
        )
    }
}