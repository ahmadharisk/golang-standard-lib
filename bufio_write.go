package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("halo dunia\n")
	_, _ = writer.WriteString("selamat\n")
	_ = writer.Flush()
}
