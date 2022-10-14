```Go
func breakPalindrome(palindrome string) string {
    s := []rune(palindrome)
    if len(s) == 1 {
        return ""
    }
    for i := 0; i < len(s)/2; i++ {
        if s[i] != 'a' {
            s[i] = 'a'
            return string(s)
        }
    }
    s[len(s)-1] = 'b'
    return string(s)
}
```