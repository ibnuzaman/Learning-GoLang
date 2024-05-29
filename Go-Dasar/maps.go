package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person ["name"] = "ibnu"
	// person ["asal"] = "imy"

	person := map[string]string{
		"name":    "ibnu",
		"address": "imy",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Print(person)

}
