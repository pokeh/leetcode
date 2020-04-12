type MinStack struct {
    stack []element
}

type element struct {
    val int
    min int
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(x int)  {
    e := element {
        val: x,
        min: this.calcMin(x),
    }
    this.stack = append(this.stack, e)
}

func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].min
}

func (this *MinStack) calcMin(x int) int {
    if len(this.stack) == 0 {
        return x
    }

    prevMin := this.stack[len(this.stack)-1].min
    if x < prevMin {
        return x
    }
    return prevMin
}
