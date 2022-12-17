package main

import "fmt"

// reference : https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

func main() {
	// basic for-each loop (slice or array)
	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	// range loop iterates over Unicode code points.
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	// loop over individual bytes
	const s = "日本語"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// map iteration: keys and values
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// channel iteration
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
