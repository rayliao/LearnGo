package main

import "github.com/rayliao/learninggo/ch8"

// say func
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

func main() {
	ch8.Log()

	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	// go say("World")
	// say("Hello")
}
