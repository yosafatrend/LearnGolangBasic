package main

import "fmt"

func main() {
	//closure : scope untuk berinteraksi dengan variable disekitarnya
	name := "Yosafat"
	counter := 0

	increment := func() {
		name = "Rendi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
