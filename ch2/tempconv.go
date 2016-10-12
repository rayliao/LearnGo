package main

import "fmt"

// Celsius : 摄氏温度
type Celsius float64

// Fahrenheit : 华氏温度
type Fahrenheit float64

// 绝对零点，冻结点，沸点
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Println("yeah")
}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
