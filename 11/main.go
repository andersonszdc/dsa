package main

import "fmt"

func calcArea(n1 [2]int, n2 [2]int) int {
	x1, y1 := n1[0], n1[1]
	x2, y2 := n2[0], n2[1]

	area := (x2 - x1) * min(y1, y2)

	return area
}

func maxArea(height []int) int {
	var p1, p2 int = 0, len(height) - 1

	var maxArea int

	for p1 < p2 {
		area := calcArea([2]int{p1, height[p1]}, [2]int{p2, height[p2]})

		if area > maxArea {
			maxArea = area
		}

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(result)
}

func init() {
	fmt.Println("test")
}