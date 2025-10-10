func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }

    key := map[rune]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    res := []string{""} // Start with an empty string as the base case

    for _, digit := range digits {
        letters := key[digit]
        var temp []string

        for _, s := range res {
            for _, letter := range letters {
                temp = append(temp, s+string(letter))
            }
        }

        res = temp // Update the result with the new combinations
    }

    return res
}
