package main

import "fmt"

func main() {
	n := Constructor()
	n.Push(1)
	fmt.Println(n.Pop())
	fmt.Println(n.Empty())
	fmt.Println("")
}

// My first attempt
type MyStack struct {
	record []int
}

func Constructor() MyStack {
	return MyStack{
		record: nil,
	}
}

func (this *MyStack) Push(x int) {
	this.record = append(this.record, x)
}

func (this *MyStack) Pop() int {
	top := this.record[len(this.record)-1]
	this.record = this.record[:len(this.record)-1]
	return top
}

func (this *MyStack) Top() int {
	top := this.record[len(this.record)-1]
	return top
}

func (this *MyStack) Empty() bool {
	if len(this.record) > 0 {
		return false
	}
	return true
}
