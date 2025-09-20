package main

import (
	"fmt"
	"slices"
)

func itIsUpperLeft(origin []int, point []int) bool {
	if point[0] <= origin[0] && point[1] >= origin[1] {
		return true
	}
	return false
}

func itIsNotOnRectangle(v1 []int, v2 []int, point []int) bool {
	if point[0] < v1[0] {
		return true
	}

	if point[0] > v2[0] {
		return true
	}

	if (point[1] > v1[1]) || (point[1] < v2[1]) {
		return true
	}

	return false
}

func numberOfPairs(points [][]int) int {
	p1Idx, p2Idx := 0, 1
	pairs := [][][]int{}

	// pairs = append(pairs, [2]int{0, 1})

	// count := 0

	for p1Idx < len(points) {
		for p2Idx < len(points) {
			p1 := points[p1Idx]
			p2 := points[p2Idx]

			if itIsUpperLeft(p1, p2) {
				pair := [][]int{p2, p1}
				pairs = append(pairs, pair)
			}

			if itIsUpperLeft(p2, p1) {
				pair := [][]int{p1, p2}
				pairs = append(pairs, pair)
			}

			p2Idx++
		}
		p1Idx++
		p2Idx = p1Idx + 1
	}

	verifiedPairs := [][][]int{}

	for _, pair := range pairs {
		isOnRectangle := false

		for _, point := range points {
			if slices.Equal(point, pair[0]) || slices.Equal(point, pair[1]) {
				continue
			}

			if !itIsNotOnRectangle(pair[0], pair[1], point) {
				isOnRectangle = true
			}
		}

		if !isOnRectangle {
			verifiedPairs = append(verifiedPairs, pair)
		}
	}

	fmt.Println(pairs)
	fmt.Println(verifiedPairs)

	return len(verifiedPairs)
}

func main() {
	numberOfPairs([][]int{{3, 1}, {1, 1}, {1, 3}})
}
