func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }

    if !wordSet[endWord] {
        return [][]string{}
    }

    // Step 1: BFS to build level map
    level := make(map[string]int)
    level[beginWord] = 0
    queue := []string{beginWord}

    for len(queue) > 0 {
        word := queue[0]
        queue = queue[1:]

        for i := 0; i < len(word); i++ {
            for c := 'a'; c <= 'z'; c++ {
                if rune(word[i]) == c {
                    continue
                }
                next := word[:i] + string(c) + word[i+1:]
                if wordSet[next] {
                    if _, seen := level[next]; !seen {
                        level[next] = level[word] + 1
                        queue = append(queue, next)
                    }
                }
            }
        }
    }

    res := [][]string{}
    if _, ok := level[endWord]; !ok {
        return res
    }

    // Step 2: DFS to build all shortest paths
    path := []string{endWord}
    dfs(endWord, beginWord, level, &res, path)

    return res
}

func dfs(word, beginWord string, level map[string]int, res *[][]string, path []string) {
    if word == beginWord {
        // Reverse the path before adding
        rev := make([]string, len(path))
        for i := range path {
            rev[len(path)-1-i] = path[i]
        }
        *res = append(*res, rev)
        return
    }

    for i := 0; i < len(word); i++ {
        for c := 'a'; c <= 'z'; c++ {
            if rune(word[i]) == c {
                continue
            }
            next := word[:i] + string(c) + word[i+1:]
            if lvl, ok := level[next]; ok && lvl == level[word]-1 {
                dfs(next, beginWord, level, res, append(path, next))
            }
        }
    }
}
