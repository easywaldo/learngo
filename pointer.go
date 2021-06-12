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

	var f *int
	failedUpdate(f)
	fmt.Println(f)

	z := 10
	failedUpdate(&z)
	fmt.Println(z)
	update(&z)
	fmt.Println(z)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int) {
	*px = 20
}
