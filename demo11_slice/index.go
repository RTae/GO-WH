package main

import "fmt"

func main() {
	var number1 = make([]int, 3, 5) // 3 len, 5 cap
	number1 = append(number1, 2)
	number1 = append(number1, 2)
	number1 = append(number1, 2)
	showSlice(number1)

	var number2 []int // slice
	number2 = append(number2, 1)
	number2 = append(number2, 2)
	showSlice(number2)
}

func showSlice(x []int) {
	fmt.Printf("Len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
