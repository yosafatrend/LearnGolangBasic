package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Yosafat"
	names[1] = "Rendi"
	names[2] = "Prayoga"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		12,
		23,
		34,
	}

	fmt.Println(values)

	fmt.Println(len(names)) //panjang array bkn jmlh data
}
