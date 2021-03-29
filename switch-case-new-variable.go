package main

import (
	"fmt"
)

func main() {
	words := []string{"cow", "cat", "dog", "bird", "a", "elephant", "snake", "kangaroo", "development"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			newLen := len(word)
			fmt.Println(word, "is exactly the right length: ", newLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}
