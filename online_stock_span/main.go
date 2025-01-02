package main

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([][2]int, 0)}
}

func (s *StockSpanner) Next(price int) int {
	ans := 1
	for len(s.stack) > 0 && s.stack[len(s.stack)-1][0] <= price {
		ans += s.stack[len(s.stack)-1][1]
		s.stack = s.stack[:len(s.stack)-1] // Pop the stack
	}
	s.stack = append(s.stack, [2]int{price, ans}) // Push the current price and span
	return ans
}

func main() {
	spanner := Constructor()
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	for _, price := range prices {
		result := spanner.Next(price)
		println(result)
	}
}
