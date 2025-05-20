package main

import "fmt"

type Speak interface {
	talk()
}
type dog struct {
}

type cat struct{}

func (d dog) talk() {
	fmt.Println("dog speak")
}
func (c cat) talk() {
	fmt.Println("cat speak")
}
 func cantalk(s Speak){
	s.talk()
 }
func main() {
 c:=cat{}
 d:=dog{}
 cantalk(c)
 cantalk(d)
}