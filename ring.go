package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(3)
	data.Value = 1
	data = data.Next()
	data.Value = 2
	data = data.Next()
	data.Value = 3

	for i := 0; i < data.Len(); i++ {
		data.Value = "nilai " + strconv.Itoa(i+0)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
