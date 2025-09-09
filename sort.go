package main

import (
	"fmt"
	"sort"
)

type Animal struct {
	name string
	age  int
}
type AnimalSlice []Animal

func (a AnimalSlice) Len() int {
	return len(a)
}

func (a AnimalSlice) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a AnimalSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	animals := []Animal{
		{"ayam", 1},
		{"bebek", 4},
		{"cendrawasih", 2},
		{"domba", 3},
	}

	sort.Sort(AnimalSlice(animals))
	fmt.Println(animals)
}
