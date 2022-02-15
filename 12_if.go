package main

import "fmt"

func main() {
	name := "Yosafat"

	if name == "Yosafat" {
		fmt.Println("Hello", name)
	} else if name == "Akanichiyu" {
		fmt.Println("Hello Aka")
	} else {
		fmt.Println("Um.. hello?")
	}
	//short statement if w/ declaration
	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah cukup")
	}
}
