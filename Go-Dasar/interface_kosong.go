// interface kosong bisa mengirim data apapun

package main

import "fmt"

func interfaceKosong() interface{} {
	return true
}

func main() {

	kosong := interfaceKosong()

	fmt.Println(kosong)
}
