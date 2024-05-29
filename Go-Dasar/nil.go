package main

// nil hanya bisa digunakan pada tipedata interface, function, map, slice, pointer, dan channel

import "fmt"

// func testTipedata(name string) string {
// 	if name == ""{
// 		return nil // eror
// 	}
// }

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{

			"name": name,
		}
	}
}

func main() {
	data := newMap("Ibnu")

	if data == nil {
		fmt.Print("Data masih kosong")
	} else {
		fmt.Println(data["name"])
	}

}
