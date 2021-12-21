package main

import (
	"fmt"
)

const chinese = "Chinese"
const chinesesHelloPrefix = "你好, "
const japan = "Japan"
const japanHelloPrefix = "こんにちは, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chinesesHelloPrefix
	case japan:
		prefix = japanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}
