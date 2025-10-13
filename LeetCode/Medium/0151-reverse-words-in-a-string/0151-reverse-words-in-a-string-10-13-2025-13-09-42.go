func reverseWords(s string) string {
    words := strings.Split(s," ")
    n := len(words)
    clean := []string{}
    for i := n-1; i>=0; i-- {
        if words[i] == "" {
            continue
        }
        clean = append(clean,words[i])
    }
    return strings.Join(clean," ")
}