func removeDuplicates(s string) string {
    // s := s
    i := 0
    for i < len(s)-1 {
        if s[i] == s[i+1]{
            if i == 0 {
                if i+2 < len(s) {
                    s = s[i+2:]
                }else{
                    s = ""
                }
                i=0
            }else{
                temp := s[:i]
                s = temp+s[i+2:]
                i -= 1
            }
        }else{
            i += 1
        }
        // fmt.Println(s)
    }
    return s
}