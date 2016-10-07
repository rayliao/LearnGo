package main

import (
	"fmt"
)

const (
	x = iota
	y
	z
	w
)

func main() {
	var c complex64 = 5 + 5i
	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)
	fmt.Print(w)
	fmt.Printf("Value is: %v", c)
}
