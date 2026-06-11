package main

import (
	"fmt"
	"sync"
)

func main() {
	// go say("hello")
	// say("world")

	// say("hello")
	// go say("world")

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go say("hello", wg)
	go say("world", wg)

	wg.Wait()
}

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(s)
	}
}
