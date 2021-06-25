package main

import (
	"fmt"
	"strconv"
	"strings"
)

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

func searchFunc(literal string) []string {
	slice := strings.Split(literal, " ")
	return slice
}

func indexWord(input string) []string {
	temp := strings.Split(input, " ")
	result := []string{}
	for index, word := range temp {
		//fmt.Println(index)
		//fmt.Println(word)
		//result += strconv.Itoa(index) + ":" + word
		result = append(result, strconv.Itoa(index)+":"+word)
	}
	fmt.Println("result is", result)
	return result
}

func toUpper(input string) []string {
	return searchFunc(strings.ToUpper(input))
}

func main() {
	searchResult := searchData("easy waldo", []func(string) []string{toUpper, indexWord})
	fmt.Println(searchResult)
	//indexWord([]string{"easy", "waldo"})
}
