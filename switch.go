package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(100); {
	case n == 0:
		fmt.Println("it's too low : ", n)
	case n > 5:
		fmt.Println("it's too big :", n)
	default:
		fmt.Println("it's good number :", n)
	}
}
