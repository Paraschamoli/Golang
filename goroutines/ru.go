package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("task",id);
}
func main() {
for i:=0;i<11;i++{
	go task(i)

	go func ()  {
		fmt.Println("hello")
	}()
}

time.Sleep(time.Second*2)
}