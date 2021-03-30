package main

import (
	"errors"
	"fmt"
)

func main() {
	firstName := MyFunc(MyFuncOpts{
		FirstName: "lee",
	})
	fmt.Println(firstName)

	resultNum := addTo(10, []int{1, 1, 1, 1, 1, 1, 1, 1}...)
	fmt.Println(resultNum)

	result, remainder, err := divAndRemainder(10, 0)
	fmt.Println(result, remainder, err)

	result, remainder, err = divAndRemainder(10, 2)
	fmt.Println(result, remainder, err)
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

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}
