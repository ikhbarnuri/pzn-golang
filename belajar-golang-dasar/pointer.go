package main

import "fmt"

type address struct {
	City, Province, Country string
}

func main() {
	address1 := address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1
	//address2.City = "Sukabumi"
	//
	//fmt.Println(address1)
	//fmt.Println(address2)

	address2 := &address1
	address2.City = "Sukabumi"

	fmt.Println(address1)
	fmt.Println(address2)
}
