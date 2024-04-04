func maxDepth(s string) int {
	max := 0
    count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
			if count > max {
				max = count
			}
		} else if ch == ')' {
			count--       
		}
	}
	return max
}