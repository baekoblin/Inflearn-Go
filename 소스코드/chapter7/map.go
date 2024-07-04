package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	delete(m, "two")
	fmt.Println(m)

	value, exists := m["one"]

	fmt.Println(value, exists)

	nestedMap := map[string]map[string]int{
		"group1": {"one": 1, "two": 2},
		"group2": {"three": 3, "four": 4},
	}

	fmt.Println(nestedMap)
	fmt.Println(len(nestedMap))
	fmt.Println(nestedMap["group3"])

	m["four"] = 4

	fmt.Println(m["four"])
}
