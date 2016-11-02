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

	// go say("World")
	// say("Hello")
}
