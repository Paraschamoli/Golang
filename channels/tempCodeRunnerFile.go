package main

import (
	"fmt"
	"time"
)

func sum(ch chan int, num1, num2 int) {
	fmt.Println("before sum")
	result:=num1+num2
	ch <- result
	fmt.Println("after sum")
	for {
		fmt.Println(22)
		time.Sleep(time.Nanosecond)
	}
}
func main() {
	ch := make(chan int)

	go sum(ch,6,4)
	result:=<-ch

	fmt.Println("the sum is =",result)

}