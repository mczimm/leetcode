package main

import "fmt"

func main() {
	n := Constructor()
	n.Push(1)
	n.Push(2)
	fmt.Println(n.Peek())
	fmt.Println(n.Pop())
	fmt.Println(n.Empty()) //[null, null, null, 1, 1, false]
}

// My first attempt
type MyQueue struct {
	record []int
}

func Constructor() MyQueue {
	return MyQueue{
		record: nil,
	}
}

func (this *MyQueue) Push(x int) {
	this.record = append(this.record, x)
}

func (this *MyQueue) Pop() int {
	if len(this.record) > 0 {
		r := this.record[0]
		this.record = this.record[1:]
		return r
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if len(this.record) > 0 {
		return this.record[0]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return len(this.record) <= 0
}
