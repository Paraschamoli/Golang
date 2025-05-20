package main

import "fmt"

func add(a int , b int) int {
	return a + b
}
func divide(a int, b int) (int, int) {
    return a / b, a % b
}

func first(na func(int,int) int)int{
	return na(6,7)
}
func main() {
	fmt.Println(add(4,6))
	ans,rem:=divide(4,6)
	fmt.Println(ans,rem)

	fmt.Println(first(add))
}