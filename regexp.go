package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`pe([a-z])po`)

	fmt.Println(regex.MatchString("peppo"))
}
