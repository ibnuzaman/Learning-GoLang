package main

// pass by reference

import "fmt"

type Addres struct {
	City, Prov, Country string
}

func main() {

	// pas by value
	// var address1 Addres = Addres{"Indramayu", "Jawa Barat", "Indonesia"}
	// var address2 = address1
	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// pas by reference

	var address1 Addres = Addres{"Indramayu", "Jawa Barat", "Indonesia"}
	var address2 *Addres = &address1
	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}
