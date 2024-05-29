package main

import "fmt"

// interface
type HasName interface {
	GetName() string
}

// fungsi
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

//struct & implementasi
type Person struct {
	Name string
}

//implementasi
func (person Person) GetName() string {
	return person.Name
}

// contoh lagi
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

//fungsi main
func main() {
	person := Person{Name: "Kennedy"}
	SayHello(person)

	animal := Animal{Name: "Mamalia"}
	SayHello(animal)
}
