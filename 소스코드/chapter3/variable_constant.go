package main

import "fmt"

type Person struct {
	Name  string
	Phone int
}

type Speaker interface {
	Speak() string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Dog struct{}

func main() {
	const PI = 3.14
	const GREETING = "Hello, World"
	fmt.Println(PI)
	fmt.Println(GREETING)

	const (
		MONDAY  = 1
		TUESDAY = 2
	)
	fmt.Println(MONDAY)
	// Pi = 3.15 error

	const (
		A = iota
		B
		C
	)
	fmt.Println(C)

	var name string = "John"
	fmt.Println(name)

	// name := "baekoblin" -> 이미 선언된 변수이니 := 는 사용 불가능합니다!
	name2 := "baekoblin"
	fmt.Println(name2)

	var (
		width, height int = 100, 200
		sumation          = 300
	)

	fmt.Println(width, height, sumation)

	var noinitValue int
	fmt.Println(noinitValue)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice []int = []int{5, 6, 7, 8, 9, 10}
	var m map[string]int = map[string]int{"map": 1, "check": 2}
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(m)

	var bob Person = Person{Name: "Bob", Phone: 123123123}
	fmt.Println(bob)

	var x int = 10
	var y *int = &x
	fmt.Println(x, y)

	var happy Dog = Dog{}
	fmt.Println(happy.Speak())

}
