// type MinStack struct {
//     stack [][]int
// }


// func Constructor() MinStack {
//     return MinStack{
//         stack : [][]int{},
//     }
// }


// func (this *MinStack) Push(val int)  {
//     n := len(this.stack)
//     if n > 0 {
//         if val < this.stack[n-1][1] {
//             this.stack = append(this.stack,[]int{val,val})
//             return
//         }
//         this.stack = append(this.stack,[]int{val,this.stack[n-1][1]})
//         return
//     }
//     this.stack = append(this.stack,[]int{val,val})
// }


// func (this *MinStack) Pop()  {
//     n := len(this.stack)
//     if n > 0 {
//         this.stack = this.stack[0:n-1]
//     }
// }


// func (this *MinStack) Top() int {
//     n := len(this.stack)
//     return this.stack[n-1][0]
// }


// func (this *MinStack) GetMin() int {
//     n := len(this.stack)
//     return this.stack[n-1][1]
// }


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

 type MinStack struct {
    stack []int64
    min int64
}


func Constructor() MinStack {
    return MinStack{
        stack : []int64{},
        min : 0,
    }
}


func (this *MinStack) Push(val int)  {
    n := len(this.stack)
    if n == 0 {
        this.stack = append(this.stack,int64(val))
        this.min = int64(val)
        return
    }
    if int64(val) < this.min {
        this.stack = append(this.stack,2*int64(val)-this.min)
        this.min = int64(val)
        return
    }
    this.stack = append(this.stack,int64(val))
    // fmt.Println(this.stack)
}


func (this *MinStack) Pop()  {
    n := len(this.stack)
    if n > 0 {
            if this.stack[n-1] < this.min {
                this.min = 2*this.min-this.stack[n-1]
            }
        this.stack = this.stack[0:n-1]
    }
}


func (this *MinStack) Top() int {
    n := len(this.stack)
    if n == 0 {
        return -1
    }
    // fmt.Println(this.stack)
    if this.stack[n-1] < this.min {
        return int(this.min)
    }
    return int(this.stack[n-1])
}


func (this *MinStack) GetMin() int {
    // fmt.Println(this.stack)
    return int(this.min)
}