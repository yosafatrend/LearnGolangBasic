//pass by value : isi data nya diduplikasi
//pass by reference : isi data nya tetap di tempat
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address Address = Address{"Pati", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address //"&" menjadikan pass by reference
	var address3 *Address = &address
	//var address4 *Address = &Address{"Bandung", "Jawa Timur", "Indonesia"} //untuk membuat data dengan memori baru
	address2.City = "Salatiga"

	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //data sumber tidak akan berubah, yg berubah hanya nilai variable itu sendiri (membuat memori baru)
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //gunakan * untuk merubah sumber data

	fmt.Println(address)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address) //atau bisa gunakan new() untuk membuat data dengan memori baru
	address4.City = "Kudus"
	fmt.Println(address4)
}
