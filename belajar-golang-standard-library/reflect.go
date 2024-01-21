package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueFiled := valueType.Field(i)
		fmt.Println(valueFiled.Name, "with type", valueFiled.Type)
		fmt.Println(valueFiled.Tag.Get("required"))
		fmt.Println(valueFiled.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "", ""})

	person := Person{
		Name:    "Eko",
		Email:   "eko@mail.com",
		Address: "Jakarta",
	}
	fmt.Println(isValid(person))
}
