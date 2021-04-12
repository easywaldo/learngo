package main

import "fmt"

func main() {
	var x int64
	x = 10
	fmt.Println(x)
	y := &x
	*y = 20
	fmt.Println(x)

	var pointerToX *int64
	pointerToX = &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	var errorVariable int64
	errorVariable = *y
	fmt.Println(errorVariable)
}
