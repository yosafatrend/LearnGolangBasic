package main

import "fmt"

func main() {
	var nilai32 int32 = 127
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Yosafat"
	var e = name[0] //byte = uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
