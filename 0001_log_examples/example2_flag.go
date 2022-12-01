package main

import (
	"log"
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

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	log.Printf("%s login, age:%d", u.Name, u.Age)
}
