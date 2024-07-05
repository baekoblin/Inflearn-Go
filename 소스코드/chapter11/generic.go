package main

import (
	"fmt"
)

// 제네릭 함수
func Print[T any](t T) {
	fmt.Println(t)
}

// Comparable 제약조건
func Compare[T comparable](a, b T) bool {
	return a == b
}

// 사용자 정의 제약조건
type Number interface {
	int | int64 | float64
}

// 제네릭 함수 - 사용자 정의 제약조건
func Sum[T Number](a, b T) T {
	return a + b
}

// 제네릭 타입
type Pair[T any] struct {
	First  T
	Second T
}

// 제네릭 타입 메서드
func (p Pair[T]) String() string {
	return fmt.Sprintf("Pair(%v, %v)", p.First, p.Second)
}

// 제네릭 인터페이스
type Stringer[T any] interface {
	String() string
}

// 타입 제약 및 타입 변환
func ConvertAndPrint[T, U any](value T, converter func(T) U) {
	convertedValue := converter(value)
	fmt.Println(convertedValue)
}

func main() {
	// 제네릭 함수 사용
	Print(42)
	Print("Hello, Go!")
	Print(3.14)

	// Comparable 제약조건 사용
	fmt.Println(Compare(5, 5))       // true
	fmt.Println(Compare("Go", "Go")) // true
	fmt.Println(Compare(3.14, 2.71)) // false

	// 사용자 정의 제약조건 사용
	fmt.Println(Sum(10, 20))                 // 30
	fmt.Println(Sum(10.5, 20.5))             // 31
	fmt.Println(Sum(int64(100), int64(200))) // 300

	// 제네릭 타입 사용
	pair := Pair[int]{First: 1, Second: 2}
	fmt.Println(pair.String())

	// 제네릭 인터페이스 사용
	var stringer Stringer[int] = Pair[int]{First: 1, Second: 2}
	fmt.Println(stringer.String())

	// 타입 제약 및 타입 변환
	ConvertAndPrint(42, func(i int) string {
		return fmt.Sprintf("Converted number: %d", i)
	})
}
