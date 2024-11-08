type MinStack struct {

    s []int
    min []int
    top int
    
}


func Constructor() MinStack {
    return MinStack{
        s : make([]int,0),
        min : make([]int,0),
        top : -1,
    }
}


func (this *MinStack) Push(val int)  {
    if this.top == -1 {
        this.min = append(this.min,val)
    } else {
        this.min = append(this.min , min(this.min[this.top],val))
    }
    this.top++
    this.s = append(this.s,val)
    
}


func (this *MinStack) Pop()  {
   
   this.s = this.s[:this.top]
   this.min = this.min[:this.top]
   this.top--

}


func (this *MinStack) Top() int {
   
    return this.s[this.top]

}


func (this *MinStack) GetMin() int {
    
   return this.min[this.top]

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */