package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", my name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30
	fmt.Println(eko)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     20,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Sukabumi", 35}
	fmt.Println(budi)

	eko.sayHello(budi.Name)
}
