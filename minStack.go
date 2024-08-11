package main

type stackNode struct {
	data    int
	lastMin int
}
type MinStack struct {
	stack []*stackNode
	min   int
}

func Constructor() MinStack {
	stack := MinStack{}
	return stack
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, &stackNode{data: val, lastMin: val})
		return
	}
	if this.stack[len(this.stack)-1].lastMin > val {
		this.stack = append(this.stack, &stackNode{data: val, lastMin: val})
	} else {
		this.stack = append(this.stack, &stackNode{data: val, lastMin: this.stack[len(this.stack)-1].lastMin})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].data
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].lastMin
}
