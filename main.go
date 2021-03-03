package main

import (
	"fmt"
	"math/cmplx"
)

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
	fmt.Println(cmplx.Abs(x))

	var a, b, c int = 10, 20, 30
	fmt.Println(a, b, c)

	var (
		age  int    = 40
		name string = "waldo"
	)
	fmt.Println(age)
	fmt.Println(name)

	var weight, nickName = 70, "turbo"
	fmt.Println(weight)
	fmt.Println(nickName)

}
