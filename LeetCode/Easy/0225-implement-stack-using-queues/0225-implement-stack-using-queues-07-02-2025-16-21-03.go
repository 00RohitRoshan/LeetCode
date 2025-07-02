type MyStack struct {
    queue []int
}


func Constructor() MyStack {
    return MyStack {
        queue: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
    if len(this.queue) > 0 {
        evicted := this.queue[len(this.queue) - 1]
        this.queue = this.queue[:len(this.queue) - 1]
        return evicted
    }

    return -1
}


func (this *MyStack) Top() int {
    if len(this.queue) > 0 {
        return this.queue[len(this.queue) - 1]
    }

    return -1
}


func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */