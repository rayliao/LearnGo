package demo

import "fmt"

var name = "rayliao"

// money msg
func money(msg string) {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
}
