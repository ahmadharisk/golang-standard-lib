package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//createNewFile("sample.log", "this is sample log")
	addToFile("sample.log", "\npesan tambahan lagi")
	result, _ := readFile("sample.log")
	fmt.Println(result)
}

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
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		message += string(line) + "\n"
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	return nil
}
