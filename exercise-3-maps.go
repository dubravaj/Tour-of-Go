package main

import (
    "strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	
	wcMap := make(map[string]int)
	words := strings.Fields(s)
	
	for _, word := range(words) {
		wcMap[word] += 1
	}
	
	return wcMap
}

func main() {
	wc.Test(WordCount)
}
