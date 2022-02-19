package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Alamat) {
	address.Country = "Indonesia"
}

func main() {
	var address Alamat = Alamat{
		City:     "Pati",
		Province: "Jawa Tengah",
		Country:  "",
	}

	ChangeCountryToIndonesia(&address) //agar data yg dikirim berubah, gunakan pointer
	fmt.Println(address)
}
