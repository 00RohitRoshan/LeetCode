func updateMatrix(mat [][]int) [][]int {
    q := [][]int{}
    vis := make([][]bool,len(mat))
    dis := make([][]int,len(mat))
    for i,row := range mat {
        vis[i] = make([]bool,len(row))
        dis[i] = make([]int,len(row))
        for j,e := range row {
            if e == 0 {
                q = append(q,[]int{i,j,0})
                vis[i][j] = true
            }
        }
    }
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    l := 1
    for len(q) > 0{
        n := len(q)
        m := n
        for n > 0 {
            x,y := q[n-1][0],q[n-1][1]
            for _,dir := range dirs {
                nx,ny := x+dir[0],y+dir[1]
                if nx >= 0 && nx < len(vis) && ny >= 0 && ny < len(vis[0]) && !vis[nx][ny] {
                    vis[nx][ny] = true
                    q = append(q,[]int{nx,ny})
                    dis[nx][ny] = l
                }
            }
            n--
        }
        l++
        q = q[m:]
    }
    return dis
}