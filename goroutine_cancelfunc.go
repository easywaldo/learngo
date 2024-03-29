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

func main() {
	ch, cancel := countTo(100)
	for i := range ch {
		if i > 50 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
