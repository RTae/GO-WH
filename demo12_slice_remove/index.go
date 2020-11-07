package main

import "fmt"

func main() {
	var number1 = []int{1, 2, 3, 4, 5, 6}
	showSlice(number1)

	// leading remove
	number1 = number1[1:len(number1)]
	showSlice(number1)
	// leading remove
	number1 = number1[1:len(number1)]
	showSlice(number1)
	// leading remove
	number1 = number1[1:len(number1)]
	showSlice(number1)

	//trailling remove
	number1 = number1[0 : len(number1)-1]
	showSlice(number1)
	//trailling remove
	number1 = number1[0 : len(number1)-1]
	showSlice(number1)

	var number = []int{1, 2, 3, 4, 5, 6}
	number = removeIndex(number, 2)
	showSlice(number)

}

func showSlice(x []int) {
	fmt.Printf("Len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
