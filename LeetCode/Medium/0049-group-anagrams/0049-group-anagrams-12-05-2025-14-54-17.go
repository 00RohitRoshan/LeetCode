func groupAnagrams(strs []string) [][]string {
    anaGrams := make(map[string][]int)
    for i,s := range strs {
        key := SortString(s)
        if anaGrams[key] == nil {
            anaGrams[key] = []int{i}
        }else{
            anaGrams[key] = append(anaGrams[key],i)
        }
    }
    res := make([][]string,len(anaGrams))
    i := 0
    for _,ids := range anaGrams {
        a := make([]string,len(ids))
        for i,id := range ids {
            a[i] = strs[id]
        }
        res[i] = a
        i++
    }
    return res
}

func SortString(s string) string {
	// Convert string to a slice of runes to handle Unicode characters correctly
	runes := []rune(s)

	// Use sort.Slice to sort the rune slice in-place
	// The provided function defines the "less than" condition for two indices i and j
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j] // Sort in ascending order
	})

	// Convert the sorted rune slice back to a string
	return string(runes)
}