package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var num int = 0
	num += 100
	fmt.Println(num)
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
}
