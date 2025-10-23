func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    a := newInterval[0]
    b := newInterval[1]
    res := [][]int{}
    i := 0
    for i < n && intervals[i][1]<a {
        res = append(res,intervals[i])
        i++
    }
    for i < n && intervals[i][0]<=b {
        a = min(a,intervals[i][0])
        b = max(b,intervals[i][1])
        i++
    }
    res = append(res,[]int{a,b})
    res = append(res,intervals[i:]...)
    return res
}