func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    oldColor := image[sr][sc]
    if oldColor == color {
        return image // nothing to change
    }

    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    q := [][]int{{sr, sc}}
    image[sr][sc] = color

    for len(q) > 0 {
        point := q[0]
        q = q[1:]
        for _, d := range dirs {
            x, y := point[0]+d[0], point[1]+d[1]
            if x >= 0 && x < len(image) && y >= 0 && y < len(image[0]) && image[x][y] == oldColor {
                image[x][y] = color
                q = append(q, []int{x, y})
            }
        }
    }
    return image
}