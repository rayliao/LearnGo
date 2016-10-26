package main

import (
	"fmt"
	// "github.com/rayliao/go-learning/ch1"
	"github.com/rayliao/go-learning/ch5"
)

func main() {
	// ch1.Fetch()
	ch5.FindLinks()

	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range ch5.TopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
