package main

import (
	"fmt"
	"github.com/rayliao/LearningGo/web"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {

	web.Server()

	web.Client()

	varlues := []int{1, 2, 1}
	fmt.Println(sum(varlues...))
}
