package main

import (
	"fmt"
	"os"
	// "github.com/rayliao/go-learning/demo"
	// "github.com/rayliao/go-learning/utils"
	"github.com/rayliao/go-learning/modals"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := modals.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
	// demo.Log("hahha")
	// demo.money("为什么这样就先不")
	// utils.Log("User Info:")
}
