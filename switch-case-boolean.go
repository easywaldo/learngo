package main

import (
	"fmt"
)

func main() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3")
		case i%7 == 0:
			fmt.Println("exit the loop")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	n := 10
	switch n {
	case 3:
		fmt.Println("n is 3")
	case 10:
		fmt.Println("n is 10")
	default:
		fmt.Println("n is default")
	}
}
