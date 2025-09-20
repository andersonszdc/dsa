package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")

	if (brokenLetters == "") {
		return len(words)
	}

	canBeTyped := 0

	for _, word := range words {
		if !strings.ContainsAny(word, brokenLetters) {
			canBeTyped++
		}
	}

	return canBeTyped
}

func main() {
	fmt.Println(canBeTypedWords("abc de", ""))
}
