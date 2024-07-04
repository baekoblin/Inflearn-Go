package main

import "fmt"

func main() {
	// var arr [5]int
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	slice2 := make([]int, 5, 10)

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	copy(slice2, slice1)

	fmt.Println(slice2)

	fmt.Println(slice2[1:3])

	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)

}
