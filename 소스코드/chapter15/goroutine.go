package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)

		}
	}
}

func worker2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Work done"
}

func sayHello() {
	fmt.Println("Hello, World")
}

func sayBye(ch chan string) {
	ch <- "Good bye"
}

func saySync(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello, World")
}
func main() {
	go sayHello()
	time.Sleep(1 * time.Second)

	go func() {
		fmt.Println("Hello from anonymous goroutine!")
	}()

	time.Sleep(1 * time.Second)

	ch := make(chan string)
	go sayBye(ch)
	msg := <-ch
	fmt.Println(msg)

	var wg sync.WaitGroup
	wg.Add(1)
	go saySync(&wg)
	wg.Wait()

	stop := make(chan bool)
	go worker(stop)
	time.Sleep(2 * time.Second)
	stop <- true
	time.Sleep(1 * time.Second)

	ch = make(chan string)
	go worker2(ch)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
