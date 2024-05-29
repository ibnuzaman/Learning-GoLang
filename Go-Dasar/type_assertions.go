package main

import (
	"fmt"
)

// bisa merubah tipe data sesuai keinginan
// biasa berpasangan dengan interface kosong

// bisa terjadi panic ketika tipe data tidak diketahui, solve nya adalah menggunakan switch

// any == interface{}

func Random() interface{} {
	return 100
}

func main() {
	var result any = Random()

	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Print(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Print("string: ", value)
	case int:
		fmt.Print("int: ", value)
	case bool:
		fmt.Print("bool: ", value)

	}
}
