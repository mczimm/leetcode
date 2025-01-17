package main

import "sort"

func main() {
	println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})) //2
	println(minMeetingRooms([][]int{{13, 15}, {1, 13}}))          //1
	println(minMeetingRooms([][]int{{9, 10}, {4, 9}, {4, 17}}))   //2
}

func minMeetingRooms(intervals [][]int) int {
	l := len(intervals)
	start := make([]int, l)
	end := make([]int, l)
	var res, j int

	for i := 0; i < l; i++ {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}

	sort.Ints(start)
	sort.Ints(end)

	for i := 0; i < l; i++ {
		if start[i] < end[j] {
			res++
		} else {
			j++
		}
	}
	return res
}
