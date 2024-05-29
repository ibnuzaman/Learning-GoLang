package main

import "fmt"

type Addres struct {
	Country string
}

func NewAddres(addres *Addres) {
	addres.Country = "Indonesia"
}

func main() {
	var address1 Addres = Addres{}

	NewAddres(&address1)
	fmt.Print(address1)
}
