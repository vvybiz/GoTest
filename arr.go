package main

import "fmt"

func main() {

	todoList := [4]string{
		"11", "22", "33", "44",
	}

	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item)
	}

}
