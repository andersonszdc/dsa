package main

import (
	"fmt"
	"slices"
	"strings"
)

func sortVowels(s string) string {
	vowels := []string{"a", "e", "i", "o", "u"}

	var allVowels []rune
	runes := []rune(s)
	vowelIdx := 0

	for _, val := range runes {
		if slices.Contains(vowels, strings.ToLower(string(val))) {
			allVowels = append(allVowels, val)
		}
	}

	slices.Sort(allVowels)

	for idx, val := range runes {
		if slices.Contains(vowels, strings.ToLower(string(val))) {
			runes[idx] = allVowels[vowelIdx]
			vowelIdx++
		}
	}

	return string(runes)
}

func main() {
	fmt.Println("lEetcOde")
	fmt.Println(sortVowels("lEetcOde"))
}
