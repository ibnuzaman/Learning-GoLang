package main

import "fmt"

type Addres struct {
	City, Prov, Country string
}

func main() {
	var address1 Addres = Addres{"Indramayu", "Jawa Barat", "Indonesia"}
	var address2 *Addres = &address1
	address2.City = "Bandung"

	*address2 = Addres{"Banten", "Banten", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
