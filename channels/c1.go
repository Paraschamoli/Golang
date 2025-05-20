package main

import "fmt"

func main() {
	messageChan := make(chan string)
// You try to give a gift to someone → but nobody’s there yet.
	messageChan <- "sending into channel"
//So you stand there waiting forever.
	msg := <-messageChan
//You never get to walk to the next step, which was to receive the gift.
	fmt.Println(msg)
} 