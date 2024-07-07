package main

import (
	"fmt"
	"reflect"
)

func reflectionTypeExample() {
	var x int = 42
	t := reflect.TypeOf(x)
	fmt.Println("Type:", t) // Type: int

	var y string = "hello"
	t = reflect.TypeOf(y)
	fmt.Println("Type:", t) // Type: string
}

func reflectionValueExample() {
	var x int = 42
	v := reflect.ValueOf(x)
	fmt.Println("Value:", v)       // Value: 42
	fmt.Println("Kind:", v.Kind()) // Kind: int

	var y string = "hello"
	v = reflect.ValueOf(y)
	fmt.Println("Value:", v)       // Value: hello
	fmt.Println("Kind:", v.Kind()) // Kind: string
}

func printTypeAndValue(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Printf("Type: %s, Value: %v\n", t, v)
}

func dynamicTypeExample() {
	printTypeAndValue(42)      // Type: int, Value: 42
	printTypeAndValue("hello") // Type: string, Value: hello
	printTypeAndValue(3.14)    // Type: float64, Value: 3.14
}

type Person struct {
	Name string
	Age  int
}

func printFields(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field: %s, Value: %v\n", t.Field(i).Name, field.Interface())
	}
}

func dynamicFieldAccessExample() {
	p := Person{Name: "Alice", Age: 30}
	printFields(p)
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func callMethod(i interface{}, methodName string) {
	v := reflect.ValueOf(i)
	method := v.MethodByName(methodName)
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Printf("Method %s not found\n", methodName)
	}
}

func dynamicMethodCallExample() {
	p := Person{Name: "Alice", Age: 30}
	callMethod(p, "Greet")
}

func setField(i interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(i).Elem()
	field := v.FieldByName(fieldName)
	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(value))
	} else {
		fmt.Printf("Cannot set field %s\n", fieldName)
	}
}

func dynamicValueSettingExample() {
	p := &Person{Name: "Alice", Age: 30}
	setField(p, "Name", "Bob")
	setField(p, "Age", 25)
	fmt.Printf("Updated Person: %+v\n", p)
}

type Speaker interface {
	Speak()
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func implementsInterface(i interface{}, iface interface{}) bool {
	t := reflect.TypeOf(i)
	ifaceType := reflect.TypeOf(iface).Elem()
	return t.Implements(ifaceType)
}

func interfaceImplementationCheckExample() {
	p := Person{Name: "Alice"}
	fmt.Println("Implements Speaker:", implementsInterface(p, (*Speaker)(nil))) // true
}

func main() {
	fmt.Println("=== Reflection Type Example ===")
	reflectionTypeExample()

	fmt.Println("=== Reflection Value Example ===")
	reflectionValueExample()

	fmt.Println("=== Dynamic Type Example ===")
	dynamicTypeExample()

	fmt.Println("=== Dynamic Field Access Example ===")
	dynamicFieldAccessExample()

	fmt.Println("=== Dynamic Method Call Example ===")
	dynamicMethodCallExample()

	fmt.Println("=== Dynamic Value Setting Example ===")
	dynamicValueSettingExample()

	fmt.Println("=== Interface Implementation Check Example ===")
	interfaceImplementationCheckExample()
}
