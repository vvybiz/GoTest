package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var mas = []int{}
	n := 100

	//Создаем целочисленные случайные элементы массива в диапазоне 0-99
	for i := 0; i < n; i++ {
		mas = append(mas, rand.Intn(100))
	}

	fmt.Println("Массив без сортировки:")
	fmt.Println(mas)

	sort.Ints(mas) //Сортировка

	fmt.Println("Массив после сортировки:")
	fmt.Println(mas)
}
