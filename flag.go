package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
