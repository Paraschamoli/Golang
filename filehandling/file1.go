package main

import (
	"fmt"
	"os"
)

func main() {

	data,err:=os.Open("example.txt")
	if err!=nil{
		panic(err)
	}
   defer data.Close()
	info,err:=data.Stat()
	if err!=nil{
		panic(err)
	}
	fmt.Println(info.Name())
	fmt.Println(info.IsDir())
	fmt.Println(info.Size())
	fmt.Println(info.Mode())
}