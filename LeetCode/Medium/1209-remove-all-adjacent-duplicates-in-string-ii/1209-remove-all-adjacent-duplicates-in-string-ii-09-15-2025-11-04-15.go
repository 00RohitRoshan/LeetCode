func removeDuplicates(s string, k int) string {
    stack := []rune{'#'}
    countStack := []int{0}
    for _,v := range s {
        n := len(stack)
        // fmt.Println(stack,countStack)
        if stack[n-1] == v {
            countStack = append(countStack, countStack[n-1]+1)
        }else{
            countStack = append(countStack, 1)
        }
        stack = append(stack,v)
        if k == countStack[n] {
            countStack = countStack[:n-k+1]
            stack = stack[:n-k+1]
        }
    }
    return string(stack[1:])
}