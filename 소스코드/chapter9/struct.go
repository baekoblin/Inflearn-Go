package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "Says Meow")
}

func MakeSound(s Speaker) {
	s.Speak()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	p          Person
	department string
}

func (p Person) Greeting() {
	fmt.Println("hello, i'm", p.Name)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)

	p1.Age = 31
	fmt.Println(p1.Age)

	em2 := Employee{p: Person{Name: "Bob", Age: 31}, department: "CR"}
	fmt.Println(em2)

	p1.Greeting()

	p1.HaveBirthday()
	fmt.Println(p1.Age)

	c := Cat{Name: "Whiskers"}
	MakeSound(c)
}
