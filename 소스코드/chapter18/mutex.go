// package main

// import (
// 	"fmt"

// 	"sync"
// )

// var (
// 	counter int

// 	mutex sync.Mutex
// )

// func increment(wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	mutex.Lock()

// 	counter++

// 	mutex.Unlock()

// }

// func main() {

// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {

// 		wg.Add(1)

// 		go increment(&wg)

// 	}

// 	wg.Wait()

// 	fmt.Println("Final counter value:", counter) // 1000 출력

// }

// package main

// import (
// 	"fmt"

// 	"sync"
// )

// func worker(id int, wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	fmt.Printf("Worker %d starting\n", id)

// 	// 작업 수행

// 	fmt.Printf("Worker %d done\n", id)

// }

// func main() {

// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {

// 		wg.Add(1)

// 		go worker(i, &wg)

// 	}

// 	wg.Wait()

// 	fmt.Println("All workers done")

// }

// package main

// import (
// 	"fmt"

// 	"sync"

// 	"time"
// )

// var (
// 	ready = false

// 	cond = sync.NewCond(&sync.Mutex{})
// )

// func worker(id int) {

// 	cond.L.Lock()

// 	for !ready {

// 		cond.Wait()

// 	}

// 	fmt.Printf("Worker %d starting\n", id)

// 	cond.L.Unlock()

// }

// func main() {

// 	for i := 1; i <= 3; i++ {

// 		go worker(i)

// 	}

// 	time.Sleep(1 * time.Second)

// 	cond.L.Lock()

// 	ready = true

// 	cond.L.Unlock()

// 	cond.Broadcast() // 모든 대기 중인 고루틴에 신호를 보냄

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("All workers started")

// }

// package main

// import (
// 	"fmt"

// 	"sync"
// )

// var (
// 	counter int

// 	rwMutex sync.RWMutex
// )

// func read(wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	rwMutex.RLock()

// 	fmt.Println("Read counter:", counter)

// 	rwMutex.RUnlock()

// }

// func write(wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	rwMutex.Lock()

// 	counter++

// 	rwMutex.Unlock()

// }

// func main() {

// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {

// 		wg.Add(1)

// 		go read(&wg)

// 	}

// 	for i := 0; i < 5; i++ {

// 		wg.Add(1)

// 		go write(&wg)

// 	}

// 	wg.Wait()

// 	fmt.Println("Final counter value:", counter)

// }

package main

import (
	"fmt"

	"sync"
)

var once sync.Once

func initialize() {

	fmt.Println("Initialized")

}

func worker(wg *sync.WaitGroup) {

	defer wg.Done()

	once.Do(initialize)

	fmt.Println("Worker")

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {

		wg.Add(1)

		go worker(&wg)

	}

	wg.Wait()

}
