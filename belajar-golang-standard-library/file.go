package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	message := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, err
}

func modifyFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	file.WriteString(message)
	return nil
}

func main() {
	// err := createNewFile("sample.log", "Sample log")
	// if err != nil {
	//	 return
	// }

	//result, err := readFile("sample.log")
	//if err != nil {
	//	return
	//}
	//fmt.Println(result)

	modifyFile("sample.log", "add new message")
}
