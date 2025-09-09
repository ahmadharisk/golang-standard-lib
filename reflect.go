package main

import (
	"fmt"
	"reflect"
)

type Vehicle struct {
	Name    string `Required:"true" max:"10"`
	Address string `Required:"true" max:"10"`
}

type Place struct {
	Name string `Required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println(valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println(field.Name)
		fmt.Println(field.Type)
		fmt.Println(field.Tag.Get("Required"))
		fmt.Println(field.Tag.Get("max"))
	}
}

func main() {
	readField(Vehicle{"ayam", ""})
	readField(Place{"borobudur"})
	fmt.Println(isValid(Vehicle{"ayam", "indo"}))
}

func isValid(value any) (result bool) {
	result = true

	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("Required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return true
}
