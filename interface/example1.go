package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["int"] = 123
	m["string"] = "hello"
	m["bool"] = true

	for _, v := range m {
		switch v.(type) {
		case string:
			fmt.Println(v, "is string")
		case int:
			fmt.Println(v, "is int")
		default:
			fmt.Println(v, "is other")
		}
	}
	fmt.Println(m)

}