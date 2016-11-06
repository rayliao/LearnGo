package ch8

import (
	"fmt"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

// func fibonacci(n int, c chan int) {
// 	x, y := 1, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	// 显示关闭channel
// 	close(c)
// }

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	// 这为什么是这样写for
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Log func
func Log() {
	// a := []int{7, 2, 8, -9, 0, 4}
	// c := make(chan int)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)

	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// ha := make(chan int, 2)
	// ha <- 1
	// ha <- 2
	// fmt.Println(<-ha)
	// fmt.Println(<-ha)

	// c := make(chan int, 10)
	// go fibonacci(cap(c), c)
	// // 可以不断的读取channel里面的数据
	// for i := range c {
	// 	fmt.Println(i)
	// }

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// <-c 发送c，可以这样写
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
