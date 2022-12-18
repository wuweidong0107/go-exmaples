package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalf("Error while opening file. Err: %s", err)
	}
	defer file.Close()

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileString := fileBuffer.String()

	fmt.Println(fileString)
}
