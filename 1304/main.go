package main

import "fmt"

func sumZero(n int) []int {
	result := []int{}

	for i := 1; i <= n/2; i++ {
		num1, num2 := i, -i

		result = append(result, num1, num2)
	}

	if n%2 != 0 {
		result = append(result, 0)
	}

	return result
}

func main() {
	result := sumZero(2)

	fmt.Println(result)
}
