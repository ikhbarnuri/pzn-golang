package main

import "fmt"

func random() any {
	return true
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//
	//var resutlInt int = result.(int)
	//fmt.Println(resutlInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}

}
