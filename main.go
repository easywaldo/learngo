package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
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

	countNumber := 10
	if countNumber > 1 {
		for i := 0; i < countNumber; i++ {
			fmt.Println(i)
		}

		for countNumber >= 1 {
			fmt.Println(countNumber)
			countNumber -= 1
		}

		evenValues := []int{2, 4, 6, 8, 10}
		for i, v := range evenValues {
			fmt.Println(i, v)
		}

		for _, v := range evenValues {
			fmt.Println(v)
		}

		uniqueNames := map[string]bool{"Apple": true, "Melon": true, "BuleBerry": false}
		for k := range uniqueNames {
			fmt.Println(k)
		}
	}

	// earlier go version cause map's order problem
	mapSample_1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop count ", i)
		for k, v := range mapSample_1 {
			fmt.Println(k, v)
		}
	}

	// value is copied in for-range
	evenValues := []int{4, 6, 8, 10, 12, 14}
	for _, v := range evenValues {
		v *= 2
	}
	fmt.Println(evenValues)

	// for break
	numbers := []int{1, 2, 3, 4}
	for _, v := range numbers {
		if v > 3 {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("random number : ", rand.Intn(100))
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
