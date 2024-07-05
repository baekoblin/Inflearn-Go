package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	ch3 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "Hello, World"
	}()

	select {
	case msg3 := <-ch3:
		fmt.Println(msg3)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	ch4 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		close(ch4)
	}()
	select {
	case msg, ok := <-ch4:
		if !ok {
			fmt.Println("Channel Closed")
		} else {
			fmt.Println(msg)
		}
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
