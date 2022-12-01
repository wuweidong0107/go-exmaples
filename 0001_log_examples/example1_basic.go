package main

import "log"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "Jack",
		Age:  18,
	}

	log.Printf("%s login, age:%d", u.Name, u.Age)
	log.Fatalf("Danger! hacker %s login", u.Name)
	//log.Panicf("Oh, system error when %s login", u.Name)
}
