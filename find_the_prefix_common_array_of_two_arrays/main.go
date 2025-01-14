package main

func main() {
	println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4})) //[0,2,3,4]
}

//func findThePrefixCommonArray(A []int, B []int) []int {
//	n := len(A)
//	prefixCommonArray := make([]int, n)
//	elementsInA := make([]int, 0)
//	elementsInB := make([]int, 0)
//
//	for i := 0; i < n-1; i++ {
//		elementsInA = append(elementsInA, A[i])
//		elementsInB = append(elementsInB, B[i])
//	}
//	commonCount := 0
//
//	for i := 0; i < len(elementsInA); i++ {
//		if _, ok := elementsInB[elementsInA[i]]; ok {
//			commonCount++
//		}
//		prefixCommonArray[i] = commonCount
//	}
//	return prefixCommonArray
//}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	prefixCommonArray := make([]int, n)

	seenInA := make(map[int]bool)
	seenInB := make(map[int]bool)

	commonCount := 0

	for i := 0; i < n; i++ {
		// Add elements to the seen maps
		seenInA[A[i]] = true
		seenInB[B[i]] = true

		// Check for common elements
		if seenInB[A[i]] {
			commonCount++
		}
		if seenInA[B[i]] && A[i] != B[i] {
			commonCount++
		}

		// Update the prefix common array
		prefixCommonArray[i] = commonCount
	}

	return prefixCommonArray
}
