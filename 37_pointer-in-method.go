package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() { //gunakan pointer untuk mengubah data sumber
	man.name = "Mr. " + man.name
}

func main() {
	yosafat := Man{
		name: "Yosafat",
	}

	yosafat.Married()
	fmt.Println(yosafat)
}
