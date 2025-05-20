package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		p.mu.Unlock()  
	}()
	p.mu.Lock()
	p.views += 1
	
}
func main() {
	p1 := post{
		views: 0,
	}
	var wg sync.WaitGroup
	for i:=0;i<199;i++{
		wg.Add(1)
	go p1.inc(&wg)
	}
	wg.Wait()
	fmt.Println(p1.views)
}