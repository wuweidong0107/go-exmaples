package main

import (
	"io/ioutil"
	"log"
)

func main() {
	linesToWrite := "This is an example to show how to write to file using ioutil"
	err := ioutil.WriteFile("temp_ioutil.txt", []byte(linesToWrite), 0x755)
	if err != nil {
		log.Fatal(err)
	}
}
