package main

import (
	"io"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "Jack",
		Age:  18,
	}

	writer1 := os.Stdout
	writer2, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logger := log.New(io.MultiWriter(writer1, writer2), "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
}
