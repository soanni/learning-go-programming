package main

import (
	"fmt"
	"runtime"
	"sync"
)

const MAX = 100000000
const workers = 2

func main() {
	runtime.GOMAXPROCS(2)

	values := make(chan int64, MAX)
	result := make(chan int64, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
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

	for i := 0; i < workers; i++ {
		go work()
	}

	wg.Wait()

	close(result)
	total := int64(0)
	for res := range result {
		total += res
	}

	fmt.Println("Total:", total)
}
