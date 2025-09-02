func greatestLetter(s string) string {
    m := make(map[rune]bool)
    for _,b := range s {
        m[b] = true
        // fmt.Printf("%c %d \n",b,b)
    }
    maxC := 'A'-1
    // fmt.Printf("%c \n",maxC)
    for k,_ := range m {
        _,ok := m['A'+k-'a']
        // fmt.Printf("%c %d \n",'A'+k-'a','A'+k-'a')
        if ok {
            maxC = max(maxC,'A'+k-'a')
        }
    }
    if maxC >= 'A' {
        return string(maxC)
    }
    return ""
}