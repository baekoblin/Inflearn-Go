package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func busyWork() {
	for i := 0; i < 1e7; i++ {
		_ = i * i
	}
}

func allocateMemory() {
	_ = make([]byte, 1<<20) // 1MB 할당
}

func worker() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func cpuProfiling() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	busyWork()
	time.Sleep(2 * time.Second)
}

func memoryProfiling() {
	f, err := os.Create("heap.prof")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	allocateMemory()
	pprof.WriteHeapProfile(f)
}

func goroutineProfiling() {
	f, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	go worker()
	go worker()

	time.Sleep(2 * time.Second)
	pprof.Lookup("goroutine").WriteTo(f, 0)
}

func gcTuning(gcValue int) {
	os.Setenv("GOGC", fmt.Sprint(gcValue))
	fmt.Println("GC tuning set to:", gcValue)
	busyWork()
}

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // 1KB 버퍼
	},
}

func memoryManagementOptimization() {
	buf := pool.Get().([]byte)
	defer pool.Put(buf)

	// 작업 수행
	_ = buf[:512] // 512바이트 사용
	fmt.Println("Memory management optimization completed")
}

func efficientDataStructure() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println("Slice:", slice)
}

func main() {
	profilingType := flag.String("type", "cpu", "Type of profiling: cpu, mem, goroutine, gc, memopt, datastruct")
	gcValue := flag.Int("gc", 100, "GC tuning value")
	flag.Parse()

	switch *profilingType {
	case "cpu":
		cpuProfiling()
	case "mem":
		memoryProfiling()
	case "goroutine":
		goroutineProfiling()
	case "gc":
		gcTuning(*gcValue)
	case "memopt":
		memoryManagementOptimization()
	case "datastruct":
		efficientDataStructure()
	default:
		fmt.Println("Unknown profiling type")
	}
}
