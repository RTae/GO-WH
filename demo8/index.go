package main

import (
	"fmt"
)

func main() {
	value := 10
	if value == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if value > 10 || value < 2 {
		fmt.Printf("Value > 10 || Value < 2\n")
	} else {
		fmt.Printf("Not Value > 10 || Value < 2\n")
	}

	if result := doSomething(); result == 10 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Nok")
	}
	fnSwitchCase()
}

func doSomething() int {
	return 10
}

func fnSwitchCase() {
	index := 2
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Printf("Someting else")
		break
	}
}
