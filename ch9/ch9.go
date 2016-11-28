package ch9

// var deposits = make(chan int)
// var balances = make(chan int)

// // Deposit func
// func Deposit(amount int) { deposits <- amount }

// // Balance func
// func Balance() int { return <-balances }

// func teller() {
// 	var balance int
// 	for {
// 		select {
// 		case amount := <-deposits:
// 			balance += amount
// 		case balances <- balance:
// 		}
// 	}
// }

// func init() {
// 	go teller()
// }

var balance int

// Deposit func
func Deposit(amount int) { balance += amount }

// Balance func
func Balance() int { return balance }
