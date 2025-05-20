package main

import (
	"fmt"
	
)

func sum(ch chan int, num1, num2 int) {
	fmt.Println("before sum")
	result:=num1+num2
	ch <- result
	// defer func() {ch <- result}()
	fmt.Println("after sum")
	
}
func main() {
	ch := make(chan int)

	go sum(ch,6,4)
	result:=<-ch

	fmt.Println("the sum is =",result)

}