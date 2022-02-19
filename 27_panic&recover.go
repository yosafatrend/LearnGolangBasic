package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil { //panic & recover MIRIP TRY CATCH COY
		fmt.Println("Aplikasi error dengan pesan : ", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLICATION ERROR!!")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	//panic function adalah function yg bisa kita gunakan untuk menghentikan program
	runApp(true)
	fmt.Println("Yosafat")
}
