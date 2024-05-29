package main

import "fmt"

func main() {
	name := [...]string{"iz", "az", "bz", "gz"}
	slice := name[2:4]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

}
