package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "eko" {
		return &notFoundError{"not found error"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("Validaton error:", validationErr)
		//} else if notFoundError, ok := err.(*notFoundError); ok {
		//	fmt.Println("Not Found error:", notFoundError)
		//} else {
		//	fmt.Println("Unknown error:", err.Error())
		//}

		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("Validaton error:", finalErr)
		case *notFoundError:
			fmt.Println("Not Found error:", finalErr)
		default:
			fmt.Println("Unknown error:", finalErr.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
