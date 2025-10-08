type MinStack struct {
    stack [][]int
}


func Constructor() MinStack {
    return MinStack{
        stack : [][]int{},
    }
}


func (this *MinStack) Push(val int)  {
    n := len(this.stack)
    if n > 0 {
        if val < this.stack[n-1][1] {
            this.stack = append(this.stack,[]int{val,val})
            return
        }
        this.stack = append(this.stack,[]int{val,this.stack[n-1][1]})
        return
    }
    this.stack = append(this.stack,[]int{val,val})
}


func (this *MinStack) Pop()  {
    n := len(this.stack)
    if n > 0 {
        this.stack = this.stack[0:n-1]
    }
}


func (this *MinStack) Top() int {
    n := len(this.stack)
    return this.stack[n-1][0]
}


func (this *MinStack) GetMin() int {
    n := len(this.stack)
    return this.stack[n-1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */