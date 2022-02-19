/*Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat
secara otomatis semua tipe data akan menjadi implementasi nya */
package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = Ups(1) // meskipun tau bahwa kembaliannya akan integer, tp kita harus mendeklarasikan sbg interface
	var data1 interface{} = Ups(2)
	var data2 interface{} = Ups(3)

	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(data2)
}
