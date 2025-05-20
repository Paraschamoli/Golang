package main

import "fmt"

func main() {

	var num []int

	fmt.Println(num==nil)

	var arr= make([]int,0 ,6)
	fmt.Println(len(arr), cap(arr))
	arr=append(arr, 12,3)
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
	temp:=[]int {2,3,4,5,6}
	arr=append(arr,temp...)
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
	cp:=make([]int,len(arr))
	copy(cp,arr)
	fmt.Println(cp)
    var t [] int
	fmt.Println(t==nil)
	t=append(t, 3)
	fmt.Println(t)

}