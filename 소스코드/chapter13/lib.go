package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, World")
	name := "Baekoblin"
	age := 31
	fmt.Println("My Name Is", name, age)

	formattedString := fmt.Sprintf("Name : %s Age %d", name, age)
	fmt.Println(formattedString)

	var input string
	fmt.Println("Enter Someting: ")
	fmt.Scanln(&input)
	fmt.Println("Entered: ", input)

	str := "Hello, World"
	fmt.Println("Contains World: ", strings.Contains(str, "World"))

	parts := []string{"hello", "world"}
	fmt.Println("Join: ", strings.Join(parts, ", "))

	fmt.Println("Split: ", strings.Split(str, ","))
	fmt.Println("Replace", strings.Replace(str, "World", "Go", 1))
	fmt.Println("To Upper:", strings.ToUpper(str))
	fmt.Println("To Lower:", strings.ToLower(str))

	fmt.Println("Abs(-5):", math.Abs(-5))

	// 거듭제곱 계산

	fmt.Println("Pow(2, 3):", math.Pow(2, 3))

	// 제곱근 계산

	fmt.Println("Sqrt(16):", math.Sqrt(16))

	// 삼각 함수 계산

	fmt.Println("Sin(π/2):", math.Sin(math.Pi/2))

	fmt.Println("Cos(π):", math.Cos(math.Pi))

	fmt.Println("Tan(π/4):", math.Tan(math.Pi/4))

	now := time.Now()

	fmt.Println("Current time:", now)

	// 시간 포맷팅

	formattedTime := now.Format("2006-01-02 15:04:05")

	fmt.Println("Formatted time:", formattedTime)

	// 문자열을 시간으로 파싱

	parsedTime, err := time.Parse("2006-01-02", "2023-10-01")

	if err != nil {

		fmt.Println("Error parsing time:", err)

	} else {

		fmt.Println("Parsed time:", parsedTime)

	}

	// 일정 시간 동안 대기

	fmt.Println("Sleeping for 2 seconds...")

	time.Sleep(2 * time.Second)

	fmt.Println("Awake!")

	// 타이머와 틱커 사용

	timer := time.After(2 * time.Second)

	ticker := time.Tick(1 * time.Second)

	fmt.Println("Timer and Ticker started")

	for {

		select {

		case <-timer:

			fmt.Println("Timer expired")

			return

		case t := <-ticker:

			fmt.Println("Tick at", t)

		}

	}

}
