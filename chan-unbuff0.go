package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 12
	fmt.Println(<-ch)
}
