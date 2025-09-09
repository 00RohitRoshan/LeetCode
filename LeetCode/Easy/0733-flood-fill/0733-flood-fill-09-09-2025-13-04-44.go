func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    q := [][]int{{sr,sc}}
    cr := image[sr][sc]
    for len(q) > 0 {
        n := len(q)
        c:=0
        for _,point := range q {
            for _,dir := range dirs {
                x,y := point[0]+dir[0],point[1]+dir[1]
                if x >= 0 && x < len(image) && y >= 0 && y < len(image[0]) && image[x][y] == cr {
                    image[x][y] = color
                    q = append(q,[]int{x,y})
                }
            }
            c++
            if c == n {
                break
            }
        }
        q = q[n:]
    }
    return image
}