package main

import "fmt"

func main() {
	firstName := MyFunc(MyFuncOpts{
		FirstName: "lee",
	})
	fmt.Println(firstName)
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) string {
	return opts.FirstName
}
