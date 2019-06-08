package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}

	histogram := make(map[string]int)
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		words := words(data)
		for word := range words {
			histogram[word]++
		}

		for k, v := range histogram {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}()

	select {
	case <-doneCh:
		fmt.Println("Counting is done!")
	case <-time.After(time.Microsecond * 1):
		fmt.Println("Timeout")
	}

}

func words(data []string) <-chan string {
	wordsCh := make(chan string)
	go func() {
		defer close(wordsCh)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				wordsCh <- word
			}
		}
	}()
	return wordsCh
}
