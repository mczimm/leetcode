/*
Given an array of meeting time "intervals" where "intervals[i] = {start, end}",
determine if a person could aatend all meetings.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}})) //false
	fmt.Println(canAttendMeetings([][]int{{7, 10}, {2, 4}}))            //true
	fmt.Println(canAttendMeetings([][]int{{13, 15}, {1, 13}}))          //true
}

func canAttendMeetings2(n [][]int) bool {
	nSorted := bubbleSort(n)
	for i := 0; i < len(nSorted)-1; i++ {
		if nSorted[i][1] < nSorted[i+1][1] {
			continue
		} else {
			return false
		}
	}

	return true
}

func bubbleSort(n [][]int) [][]int {
	cnt := len(n)
	for cnt > 0 {
		for i := 0; i < len(n)-1; i++ {
			if n[i][0] > n[i+1][0] {
				n[i][0], n[i][1], n[i+1][0], n[+1][1] = n[i+1][0], n[i+1][1], n[i][0], n[i][1]
			}
			cnt--
		}
	}
	return n
}

// 24.07.2024
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] <= intervals[i+1][0] {
			continue
		} else {
			return false
		}
	}

	return true
}
