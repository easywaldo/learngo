package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var employee Employee
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &employee)

	if err == nil {
		fmt.Println(employee.Name)
		fmt.Println(employee.Age)
		return
	}
	fmt.Println(err)
}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
