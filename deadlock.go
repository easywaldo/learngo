package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("go......")
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)

	/*
		ch1 <- v
		v2 <- ch2
		v2 <- ch1 <- v
	*/
}
