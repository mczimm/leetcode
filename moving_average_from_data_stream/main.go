package main

type MovingAverage struct {
	size  int
	queue []int
	sum   int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: []int{},
		sum:   0,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.queue = append(m.queue, val)
	m.sum += val

	if len(m.queue) > m.size {
		m.sum -= m.queue[0]
		m.queue = m.queue[1:]
	}

	return float64(m.sum) / float64(len(m.queue))

}
