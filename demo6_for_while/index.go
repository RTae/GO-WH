package main

import (
	"fmt"
)

func main() {
	fnFor()
	fmt.Println("")
	fnWhile()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("for-index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Printf("while-index %d\n", index)
	}
}
