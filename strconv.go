package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("false")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result2, err := strconv.Atoi("1000")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)
}
