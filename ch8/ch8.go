package ch8

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

// Log func
func Log() {
	a := []int{7, 2, 8, -9, 0, 4}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	ha := make(chan int, 2)
	ha <- 1
	ha <- 2
	fmt.Println(<-ha)
	fmt.Println(<-ha)
}
