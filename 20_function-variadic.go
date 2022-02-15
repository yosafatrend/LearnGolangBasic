package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 90, 20, 39)
	fmt.Println(total)

	//slice as vararg parameter
	slice := []int{10, 20, 230, 21, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
