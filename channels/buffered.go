package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 4
	ch <- 7
     close(ch)
	fmt.Println(<-ch,<-ch)
} 