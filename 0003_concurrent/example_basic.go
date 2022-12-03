package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		f(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
