package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))     // 12
	fmt.Println(minStoneSum([]int{4, 3, 6, 7}, 3))  // 12
	fmt.Println(minStoneSum2([]int{5, 4, 9}, 2))    // 12
	fmt.Println(minStoneSum2([]int{4, 3, 6, 7}, 3)) // 12
	fmt.Println(minStoneSum3([]int{5, 4, 9}, 2))    // 12
	fmt.Println(minStoneSum3([]int{4, 3, 6, 7}, 3)) // 12
}

func minStoneSum(piles []int, k int) int {
	sortPiles(piles)
	for k > 0 {
		cur := piles[0]
		dif := cur / 2
		piles = piles[1:]
		piles = append(piles, cur-dif)
		sortPiles(piles)
		k--
	}
	var res int
	for _, pile := range piles {
		res += pile
	}
	return res
}

func sortPiles(p []int) []int {
	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})
	return p
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minStoneSum2(piles []int, k int) int {
	h := MaxHeap(piles)
	heap.Init(&h)

	for i := 0; i < k; i++ {
		heap.Push(&h, int(math.Ceil(float64(heap.Pop(&h).(int))/2)))
		//maxx := heap.Pop(&h).(int) / 2
		//heap.Push(&h, maxx)
	}
	var sum int
	for _, pile := range h {
		sum += pile
	}
	return sum
}

func minStoneSum3(piles []int, k int) int {
	freq := make(map[int]int)
	maxPile := 0
	sum := 0

	// Count frequencies and find max pile
	for _, pile := range piles {
		freq[pile]++
		sum += pile
		if pile > maxPile {
			maxPile = pile
		}
	}

	// Process k operations
	for i := 0; i < k; i++ {
		// Find current max pile
		for maxPile > 0 && freq[maxPile] == 0 {
			maxPile--
		}
		if maxPile == 0 {
			break
		}

		// Remove one from frequency
		freq[maxPile]--

		// Calculate new pile size and update
		newSize := maxPile - maxPile/2
		freq[newSize]++
		sum -= maxPile / 2
	}

	return sum
}
