package main

import (
	"fmt"
	"math/cmplx"
	"strings"
)

const c int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

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

	unVarName := "unVar"
	fmt.Println(unVarName)

	test()

	r1 := multiply(2, 5)
	totalLen, _ := lenAndUpper("waldo")

	fmt.Println(r1)
	fmt.Println(totalLen)
}

func test() {
	fmt.Println(c, idKey, nameKey, z)
}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
