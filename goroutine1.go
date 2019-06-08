package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go count(10, 50, 10)
	go count(60, 100, 10)
	go count(110, 200, 20)
	start := 0
	stop := 50
	step := 5

	go func() {
		count(start, stop, step)
	}()
	fmt.Scanln()
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
