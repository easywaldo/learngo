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

	copyX := []int{1, 2, 3, 4, 5}
	fmt.Println("copyX : ", copyX)
	copyY := make([]int, 4)
	fmt.Println("copyY : ", copyY)
	copyNum := copy(copyY, copyX)
	fmt.Println("copyY : ", copyY, " copyNum : ", copyNum)

	arrX := []int{1, 2, 3, 4}
	fmt.Println(arrX)
	num = copy(arrX[:3], arrX[1:])
	fmt.Println(arrX, num)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
	fmt.Println(f.name)

	f.name = "waldo"
	fmt.Println(f.name)
	fmt.Println(g.name)

	g.name = f.name
	fmt.Println(f.name == g.name)
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
