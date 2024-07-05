package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 예제 1: Context 활용
	example1()

	// 예제 2: Context를 통한 Goroutine 관리 및 취소
	example2()

	// 예제 3: 타임아웃 설정
	example3()

	// 예제 4: 데드라인 설정
	example4()

	// 예제 5: Context에 값 저장
	example5()
}

func example1() {
	fmt.Println("Example 1: Context 활용")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Operation completed")
		case <-ctx.Done():
			fmt.Println("Operation cancelled")
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // 컨텍스트 취소
	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context, id int, value string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			if value != "" {
				fmt.Printf("Worker %d working with value: %s\n", id, value)
			} else {
				fmt.Printf("Worker %d working\n", id)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func example2() {
	fmt.Println("Example 2: Context를 통한 Goroutine 관리 및 취소")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i, "")
	}

	time.Sleep(2 * time.Second)
	cancel() // 모든 고루틴에 취소 신호 전달
	time.Sleep(1 * time.Second)
}

func example3() {
	fmt.Println("Example 3: 타임아웃 설정")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i, "")
	}

	time.Sleep(3 * time.Second)
}

func example4() {
	fmt.Println("Example 4: 데드라인 설정")
	ctx := context.Background()
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i, "")
	}

	time.Sleep(3 * time.Second)
}

func example5() {
	fmt.Println("Example 5: Context에 값 저장")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "example value")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i, ctx.Value("key").(string))
	}

	time.Sleep(2 * time.Second)
	cancel() // 모든 고루틴에 취소 신호 전달
	time.Sleep(1 * time.Second)
}
