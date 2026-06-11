package main

import (
	"fmt"
	"sync"
)

/*
	A Useful Mental Model

	For an unbuffered channel: sender <---- handshake ----> receiver
	The send and receive must happen together.

	For a buffered channel: sender ---> mailbox ---> receiver
	The sender can drop a value into the mailbox and leave, as long as the mailbox isn't full.
*/

func main() {
	balls := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		throwBalls("red", balls)
	}()
	go func() {
		defer wg.Done()
		throwBalls("blue", balls)
	}()

	go func() {
		wg.Wait()
		close(balls)
	}()

	for color := range balls {
		fmt.Println(color, "received")
	}
}

func throwBalls(color string, balls chan<- string) {
	fmt.Printf("throwing the %s ball \n", color)
	balls <- color
}
