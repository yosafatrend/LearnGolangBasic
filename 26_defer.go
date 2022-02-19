package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // biasakan memanggil defer diatas
	fmt.Println("Run application")
	result := 10 / value //misalkan ini error
	fmt.Println("Result =", result)
	// logging() //fungsi ini tidak akan dieksekusi jika kode diatasnya error
}

func main() {
	//defer memungkinkan untuk memanggil function setelah suatu function dieksekusi tanpa memperdulikan func tersebut error atau tidak
	runApplication(10)
}
