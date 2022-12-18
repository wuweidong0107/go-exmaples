package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileBytes, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fmt.Println(fileString)
}
