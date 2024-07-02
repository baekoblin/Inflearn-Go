package main

import "fmt"

func main() {
	conditionVarA := true

	if conditionVarA {
		fmt.Println("Yes Condition is True")
	} else {
		fmt.Println(" No Condition is False")
	}

	conditionVarB := false

	if conditionVarB {
		fmt.Println("Condition B is True")
	} else if conditionVarA {
		fmt.Println("Condition A is True, Condition B is False")
	} else {
		fmt.Println("All condition False")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0

	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {

	// }

	sliceA := []int{1, 2, 3, 4, 5}
	for index, value := range sliceA {
		fmt.Println(index, value)
	}

	mapA := map[string]int{"one": 1, "two": 2}
	for key, value := range mapA {
		fmt.Println(key, value)
	}

	color := "red"

	switch color {
	case "red":
		fmt.Println("Color is Red")
		fallthrough
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color?")
	}

	var x interface{} = 10

	switch v := x.(type) {
	case int:
		fmt.Println(v, " Type is int")
	}

	for i := 0; i > 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

OuterLoop:
	for i := 0; i > 10; i++ {
		for j := 0; j > 10; i++ {
			if i == 3 && j == 2 {
				break OuterLoop //continue // goto
			}
			fmt.Println(i, j)
		}
	}
}
