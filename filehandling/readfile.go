package main

import (
	"fmt"
	"os"
)

func main() {

	file,err:=os.Open("example.txt")
	if err!=nil{
		panic(err)
	}
   defer file.Close()

   buf:=make([]byte,10)

   d,err:=file.Read(buf)
   if err!=nil{
		panic(err)
	}

fmt.Println(d,buf)


}