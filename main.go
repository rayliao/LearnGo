package main

import (
	"fmt"

	"time"

	"github.com/rayliao/LearningGo/ch9"
)

// say func
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

func main() {
	// ch8.Log()

	// ha := make(chan int, 2)
	// ha <- 1
	// ha <- 2
	// fmt.Println(<-ha)
	// fmt.Println(<-ha)

	go func() {
		ch9.Deposit(200)
		fmt.Println("a=", ch9.Balance())
	}()

	go func() {
		ch9.Deposit(100)
		fmt.Println("b=", ch9.Balance())
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("4444")

	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	// go say("World")
	// say("Hello")
}
