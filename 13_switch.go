package main

import "fmt"

func main() {
	name := "Yosafat"

	switch name {
	case "Yosafat":
		fmt.Println("Hello Yosafat")
	case "Akanichiyu":
		fmt.Println("Hello Akanichiyu")
	default:
		fmt.Println("Um.. hello?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("terlalu panjang")
	// case false:
	// 	fmt.Println("udah cukup")
	// }

	//switch without condition
	length := len(name)
	switch {
	case length > 18:
		fmt.Println("terlalu panjang")
	case length > 5:
		fmt.Println("lumayan panjang")
	default:
		fmt.Println("Nama udah bener")
	}
}
