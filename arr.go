package main

import "fmt"

func main() {

	todoList := []string{
		"11", "22", "33", "44",
	}

	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	for i, item := range reverse(todoList) {
		fmt.Printf("%d. %s\n", i+1, item)
	}

}

func reverse(todoList []string) []string {

	var reverse []string

	for i := len(todoList) - 1; i >= 0; i-- {
		reverse = append(reverse, todoList[i])
	}

	return reverse

}
