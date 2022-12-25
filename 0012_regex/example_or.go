package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("abc|xyz")
	match := sampleRegex.Match([]byte("abc"))

	fmt.Println(match)
	match = sampleRegex.Match([]byte("xyz"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("abcxyz"))
	fmt.Println(match)

	match = sampleRegex.Match([]byte("abd"))
	fmt.Println(match)
}
