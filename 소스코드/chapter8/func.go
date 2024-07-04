package main

import "fmt"

func simplePrint() {
	fmt.Println("Hello Function")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func simpleAdd(a, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func sumation(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func recr(n int) int {
	if n == 0 {
		return 1
	}
	return n + recr(n-1)
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	simplePrint()
	greet("Backoblin")
	fmt.Println(simpleAdd(1, 2))
	quotient, remainder := divide(5, 2)
	fmt.Println(quotient, remainder)
	fmt.Println(sumation(1, 2, 3, 4, 5))
	fmt.Println(recr((10)))

	goodbye := func(name string) {
		fmt.Println("Good Bye!", name)
	}

	goodbye("Backoblin")

	incrementA := counter()
	incrementB := counter()

	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementB())

	var funcs []func()

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}

}
