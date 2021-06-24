package main

import (
	"fmt"
)

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		fmt.Println("canceled")
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func process(number int) int {
	return number + 1
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func main() {
	// ch, cancel := countTo(100)
	// for i := range ch {
	// 	if i > 50 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	//cancel()

	ch2 := make(chan int)
	go func() {
		ch2 <- 100
		result := processChannel(ch2)
		fmt.Println(result)
	}()

	//cancel()
}
