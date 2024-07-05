package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("Worker %d processing job %d \n", id, job)
	}
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello World"
	}()

	msg := <-ch
	fmt.Println(msg)

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	ch3 := make(chan int)
	go func() {
		ch3 <- 42
		close(ch3)
	}()

	for value := range ch3 {
		fmt.Println(value)
	}

	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		ch4 <- "Hello from ch4"
	}()

	go func() {
		ch5 <- "Hello from ch5"
	}()

	select {
	case msg1 := <-ch4:
		fmt.Println(msg1)
	case msg2 := <-ch5:
		fmt.Println(msg2)
	}

	ch6 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch6 <- "Hello World"
	}()

	select {
	case msg := <-ch6:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	var wg sync.WaitGroup
	ch9 := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(1, &wg, ch9)
	}

	for j := 1; j <= 5; j++ {
		ch9 <- j
	}

	close(ch9)
	wg.Wait()
}
