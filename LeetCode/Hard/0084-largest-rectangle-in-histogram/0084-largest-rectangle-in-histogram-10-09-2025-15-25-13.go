func largestRectangleArea(heights []int) int {
    numberOfHeights := len(heights)
    stack := []int{}
    left := make([]int,numberOfHeights)
    right := make([]int,numberOfHeights)
    for i:= range numberOfHeights{ 
        for len(stack)>0 && heights[stack[len(stack)-1]] > heights[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            left[i] = i
        }else{
            left[i] = stack[len(stack)-1]+1
        }
        stack = append(stack,i)
    }
    stack = []int{}
    for i:=numberOfHeights-1; i>=0; i--{ 
        for len(stack)>0 && heights[stack[len(stack)-1]] > heights[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            right[i] = i
        }else{
            right[i] = stack[len(stack)-1]-1
        }
        stack = append(stack,i)
    }
    maxheight := math.MinInt
    for i := range numberOfHeights{
        maxheight = max(maxheight,(right[i]-left[i]+1)*heights[i])
    }
    return maxheight
}

// // TLE
// func largestRectangleArea(heights []int) int {
//     n := len(heights)
//     maxArea := math.MinInt
//     for i := 0; i<n; i++{
//         minHeight := math.MaxInt
//         for j := i; j<n; j++{
//             minHeight = min(minHeight,heights[j])
//             maxArea = max(maxArea,minHeight*(j-i+1))
//         }
//     }
//     return maxArea
// }