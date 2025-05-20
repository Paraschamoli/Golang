package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go task(i, &wg)

		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("hello")
		}()
	}

	wg.Wait() // Wait until all goroutines are done
	fmt.Println("All goroutines finished.")
}
