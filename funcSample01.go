package main

import (
	"fmt"
	"strings"
)

func lengthAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func main() {
	length, name := lengthAndUpper("waldo")
	fmt.Println(length, name)
}
