package main

import "fmt"

func f1() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	counter := f1()
	counter2 := f1()
	fmt.Println(counter())
	fmt.Println(counter(),counter2())
}