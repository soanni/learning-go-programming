package main

import (
	"fmt"
	"sync"
)

const MAX = 100000000

func main() {
	values := make(chan int64, MAX)
	result := make(chan int64, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				values <- int64(i)
			}
		}
		close(values)
	}()

	work := func() {
		defer wg.Done()
		var r int64
		r = 0
		for i := range values {
			r += i
		}
		result <- int64(r)
	}

	go work()
	go work()

	wg.Wait()

	total := <-result + <-result
	fmt.Println("Total:", total)
}
