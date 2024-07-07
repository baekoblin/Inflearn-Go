package main

import (
	"errors"
	"fmt"
	"os"
)

// error 인터페이스 개념
// Go언어의 에러 처리는 `error` 인터페이스를 사용
// `error` 인터페이스는 `Error()` 메서드를 포함하며, 에러 메시지를 반환

// type error interface {
//     Error() string
// }

// 기본 에러 처리 방식
// 함수는 에러를 반환할 수 있으며, 호출자는 반환된 에러를 확인하여 처리
// 에러가 발생하면 `nil`이 아닌 `error` 값을 반환

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
// `error` 인터페이스를 구현하여 사용자 정의 에러 타입을 생성
// 사용자 정의 에러 타입은 추가적인 정보를 포함할 수 있음

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
// errors.New 함수
// `errors.New` 함수를 사용하여 새로운 에러를 생성

func basicError() {
	err := errors.New("an error occurred")
	fmt.Println("Error:", err)
}

// fmt.Errorf 함수
// `fmt.Errorf` 함수를 사용하여 포맷팅된 에러 메시지를 생성

func formattedError() {
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println("Error:", err)
}

// panic 및 recover
// panic 함수의 사용법 및 사례
// `panic` 함수는 프로그램의 실행을 중단하고, 현재 고루틴의 패닉 상태를 발생
// 주로 복구할 수 없는 심각한 오류가 발생했을 때 사용

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

// recover 함수로 패닉 복구하기
// `recover` 함수는 패닉 상태를 복구하고, 프로그램의 실행을 계속
// `defer` 키워드를 사용하여 패닉 발생 시 `recover` 함수를 호출

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

// 패닉과 복구의 활용 사례
// 패닉과 복구를 사용하여 예외 상황을 처리하고, 프로그램의 안정성을 높일 수 있음

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

func main() {
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
}
