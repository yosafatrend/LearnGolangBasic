//adalah konversi tipe data dari interface kosong
package main

import "fmt"

func random() interface{} {
	return "Yosafat"
}

func main() {
	var result interface{} = random()
	//tapi ini tidak aman, jika tipe data nya salah akan menimbulkan panic
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	//ini yang aman
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
