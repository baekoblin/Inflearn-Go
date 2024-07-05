package utils

import "fmt"

func Greet() {
	fmt.Println("hello it's me")
}

func privacyGreet() {
	fmt.Println("Privacy Greet")
}

var PublicName = "baekoblin"
var privacyName = "Unknowns"

func init() {
	fmt.Println("My Utils initialized")
}
