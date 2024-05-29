package main

import "fmt"

type Addres struct {
	City, Prov, Country string
}

func main() {
	var alamat1 *Addres = new(Addres)
	alamat1.City = "Indramayu"
	alamat1.Prov = "Jawa Barat"
	alamat1.Country = "Indonesia"
	var alamat2 *Addres = alamat1
	alamat2.City = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
