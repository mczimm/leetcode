package main

import "testing"

var ret int

func BenchmarkPart1(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = shortestDistance([]string{"this", "is", "a", "long", "sentence", "is", "fun", "day", "today", "sunny", "weather", "is", "a", "day", "tuesday", "this", "sentence", "run", "running", "rainy"}, "weather", "long")
	}
	ret = r
}
