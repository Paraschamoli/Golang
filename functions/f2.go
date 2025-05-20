package main

import (
	"fmt"
)

func checkeven(arr []int, condition func(int) bool) []int {
	var temp []int
	for _, val := range arr {
		if condition(val) {
			temp = append(temp, val)
		}
		
	}
	return temp
}
func main() {
	arr := []int{1, 3, 2, 4, 5, 6, 7, 8}
	result := checkeven(arr, func(n int) bool {
		return n%2==0
	})
	fmt.Println(result)
}