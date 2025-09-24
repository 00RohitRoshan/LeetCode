func findOrder(numCourses int, prerequisites [][]int) []int {
    adj := make([][]int,numCourses)
    inDegree := make([]int,numCourses)
    for _,course := range prerequisites {
        cur,pre := course[0],course[1]
        adj[pre] = append(adj[pre],cur)
        inDegree[cur]++
    }
    leaves := []int{}
    for cur,numPre := range inDegree {
        if numPre == 0 {
            leaves = append(leaves,cur)
        }
    }
    n := numCourses
    ans := []int{}
    for n > 0{
        n -= len(leaves)
        nextLeaves := []int{}
        for _,leaf := range leaves {
            ans = append(ans,leaf)
            for _,next := range adj[leaf]{
                inDegree[next]--
                if inDegree[next] == 0 {
                    nextLeaves = append(nextLeaves,next)
                }
            }
        }
        leaves = nextLeaves
    }
    if len(ans) != numCourses {
        return []int{}
    }
    return ans
}