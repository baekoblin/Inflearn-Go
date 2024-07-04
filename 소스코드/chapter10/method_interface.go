package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello My Name is", p.Name)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

type Speacker interface {
	Speak()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "Says Meow")

}

func (d Dog) Speak() {
	fmt.Println(d.Name, "Say woof")
}

func MakeSound(s Speacker) {
	s.Speak()
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet()

	p.HaveBirthday()
	fmt.Println(p.Age)

	var s Speacker = Dog{Name: "Budy"}
	s.Speak()

	s = Cat{Name: "marry"}
	s.Speak()

	MakeSound(s)

	s = Dog{Name: "Budy"}
	MakeSound(s)

	var i interface{}

	i = 42
	fmt.Println(i)

	i = "Hello"

	fmt.Println(i)

	printType(42)
	printType(3.14)
	printType("Hello")
}
