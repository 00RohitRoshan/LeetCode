func judgeCircle(moves string) bool {
    x,y := 0,0
    for _,v := range moves{
        if v == 'U' {
            x++
        }
        if v == 'D' {
            x--
        }
        if v == 'R' {
            y++
        }
        if v == 'L' {
            y--
        }
    }

    return x==0 && y==0
}