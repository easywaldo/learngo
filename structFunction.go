package main

import "fmt"

func main() {
	firstName := MyFunc(MyFuncOpts{
		FirstName: "lee",
	})
	fmt.Println(firstName)

	resultNum := addTo(10, []int{1, 1, 1, 1, 1, 1, 1, 1}...)
	fmt.Println(resultNum)
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) string {
	return opts.FirstName
}

func addTo(base int, vals ...int) int {
	out := base
	for _, v := range vals {
		out += v
	}
	return out
}
