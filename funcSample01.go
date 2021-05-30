package main

import (
	"fmt"
	"strings"
)

func lengthAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	length, name := lengthAndUpper("waldo")
	fmt.Println(length, name)
}
