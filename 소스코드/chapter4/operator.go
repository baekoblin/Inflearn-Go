package main

import "fmt"

func main() {
	a := 5 // 0101
	b := 3 //0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

	c := true
	d := false
	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)

	e := 5
	fmt.Println(e)

	e += 5
	fmt.Println(e)

	e -= 3
	fmt.Println(e)

	e *= 2
	fmt.Println(e)

	e /= 2
	fmt.Println(e)

	e %= 3
	fmt.Println(e)

	a = 5
	b = 3
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)

	k := 10
	g := &k

	fmt.Println(*g)

}
