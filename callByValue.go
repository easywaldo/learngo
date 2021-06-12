package main

import "fmt"

func main() {
	p := person{}
	i := 20
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}

	modMap(m)
	fmt.Println(m)

	numbers := []int{2, 3, 4}
	modSlice(numbers)
	fmt.Println(numbers)
}

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "GoodBye"
	p.name = "waldo"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "good bye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
