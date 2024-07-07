package main

import (
	"errors"
	"fmt"
	"os"
)

// 에러 처리 (error 인터페이스)
// 기본 에러 처리 방식

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func basicErrorHandling() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// 사용자 정의 에러 타입

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func customDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func customErrorHandling() {
	result, err := customDivide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// errors 패키지를 이용한 에러 생성 및 처리

func basicError() {
	err := errors.New("an error occurred")
	fmt.Println("Error:", err)
}

func formattedError() {
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println("Error:", err)
}

// panic 및 recover

func panicDivide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func panicExample() {
	fmt.Println("Start")
	result := panicDivide(4, 0)
	fmt.Println("Result:", result)
	fmt.Println("End")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")
	result := panicDivide(4, 0)
	fmt.Println("Result:", result)
	fmt.Println("End")
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	result = panicDivide(a, b)
	return result, nil
}

func panicRecoveryExample() {
	result, err := safeDivide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Result:", result)
}

// go doc, go vet을 이용한 코드 문서화 및 검증

func AddDoc(a, b int) int {
	// Add 함수는 두 정수를 더한 값을 반환합니다.
	return a + b
}

func main() {
	// 에러 처리 예제
	fmt.Println("### 기본 에러 처리 방식 ###")
	basicErrorHandling()

	fmt.Println("\n### 사용자 정의 에러 타입 ###")
	customErrorHandling()

	fmt.Println("\n### errors 패키지를 이용한 에러 생성 및 처리 ###")
	basicError()
	formattedError()

	fmt.Println("\n### panic 및 recover ###")
	fmt.Println("#### panic 함수의 사용법 및 사례 ####")
	// panicExample() // Uncomment to see panic in action

	fmt.Println("#### recover 함수로 패닉 복구하기 ####")
	recoverExample()

	fmt.Println("#### 패닉과 복구의 활용 사례 ####")
	panicRecoveryExample()

	// 테스트 예제
	result := AddDoc(2, 3)
	fmt.Println("AddDoc(2, 3) =", result)
}
