package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(connectSticks([]int{1, 8, 3, 5}))  //30
	//fmt.Println(connectSticks2([]int{1, 8, 3, 5})) //30
	fmt.Println(connectSticks3([]int{1, 8, 3, 5})) //30
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	h := &MinHeap{}
	var cost int
	for _, stick := range sticks {
		heap.Push(h, stick)
	}

	for h.Len() > 1 {
		p1 := heap.Pop(h).(int)
		p2 := heap.Pop(h).(int)
		cost += p1 + p2
		heap.Push(h, p1+p2)
	}
	return cost
}

func connectSticks2(sticks []int) int {
	var cost int
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] < sticks[j]
	})
	for len(sticks) > 1 {
		cost += sticks[0] + sticks[1]
		sticks = append(sticks, sticks[0]+sticks[1])
		sticks = sticks[2:]
		sort.Slice(sticks, func(i, j int) bool {
			return sticks[i] < sticks[j]
		})
	}
	return cost
}

func connectSticks3(sticks []int) int {
	var cost int
	// Initial sort
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] < sticks[j]
	})

	for len(sticks) > 1 {
		// Take two smallest sticks
		p0 := sticks[0]
		p1 := sticks[1]
		cost += p0 + p1

		// Remove first two elements
		sticks = sticks[2:]
		// Insert new combined stick in sorted position
		newStick := p0 + p1
		i := 0
		for i < len(sticks) && sticks[i] < newStick {
			i++
		}
		sticks = append(sticks, 0)
		copy(sticks[i+1:], sticks[i:])
		sticks[i] = newStick
	}
	return cost
}
