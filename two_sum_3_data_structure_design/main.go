package main

import "fmt"

func main() {
	t := Constructor()
	t.Add(0)
	t.Add(0)
	//t.Add(1)
	fmt.Println(t.Find(0))
	//fmt.Println(t.Find(3))
	//fmt.Println(t.Find(4))
	//fmt.Println(t.Find(5))
	//fmt.Println(t.Find(6))
}

type TwoSum struct {
	sums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		map[int]int{},
	}
}

func (this *TwoSum) Add(number int) {
	this.sums[number] = number
}

func (this *TwoSum) Find(value int) bool {
	//if len(this.sums) == 2 {
	//	return this.sums[0]+this.sums[1] == value
	//}
	//if len(this.sums) < 2 {
	//	return false
	//}

	for i, v := range this.sums {
		f := value - i
		// From the solutions: && (i != f || v > 1)
		if _, ok := this.sums[f]; ok && (i != f || v > 1) {
			return true
		}
	}
	return false
}
