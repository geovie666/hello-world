package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	myMap := make(map[string]int)
	sliceOfWords:= strings.Fields(s)
	for i:=0;i<len(sliceOfWords);i++{
		counts, ok := myMap[sliceOfWords[i]]
		if ok{
			myMap[sliceOfWords[i]] = counts + 1
		} else {
			myMap[sliceOfWords[i]]=1
		}
	}
	return myMap
}

func main() {
	wc.Test(WordCount)
}
