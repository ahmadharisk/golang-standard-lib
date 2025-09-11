package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{"ayam", "bebek", "cicak", "domba"}
	values := []int{50, 60, 70, 80, 90}

	fmt.Println(slices.Contains(animals, "cicak"))
	fmt.Println(slices.Min(values))
}
