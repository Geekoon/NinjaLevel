package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	gs := 20
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			t := sum
			t++
			sum = t
			fmt.Println("add 1 in foo. sum=", sum)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("All done.", sum)
}
