package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	gs := 20
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			fmt.Println("add 1 in foo. sum=", atomic.LoadInt64(&sum))
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("i=", i, "and Sum=", atomic.LoadInt64(&sum))
		fmt.Println("CPUs\t\t", runtime.NumCPU())
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("All done.", sum)
}
