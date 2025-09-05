/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    g := make(map[int]*Employee)
    for _,j := range employees {
        g[j.Id] = j
    }
    q := []int{id}
    sum := 0
    for len(q) > 0 {
        sum += g[q[0]].Importance
        q = append(q[1:],g[q[0]].Subordinates...)
    }
    return sum
}