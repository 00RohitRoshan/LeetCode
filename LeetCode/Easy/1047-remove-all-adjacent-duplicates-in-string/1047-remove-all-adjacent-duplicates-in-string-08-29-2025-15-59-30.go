func removeDuplicates(s string) string {
    s1 := s
    i := 0
    for i < len(s1)-1 {
        if s1[i] == s1[i+1]{
            if i == 0 {
                if i+2 < len(s1) {
                    s1 = s1[i+2:]
                }else{
                    s1 = ""
                }
                i=0
            }else{
                temp := s1[:i]
                s1 = temp+s1[i+2:]
                i -= 1
            }
        }else{
            i += 1
        }
        // fmt.Println(s1)
    }
    return s1
}