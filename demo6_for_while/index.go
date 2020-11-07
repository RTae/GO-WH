package main

import (
	"fmt"
)

func main() {
	fnFor()
	fmt.Println("")
	fnWhile()
	fnWhileUsingBreak()
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
		fmt.Printf("index %d\n", index)
	}
}

func fnWhileUsingBreak() {
	index := 0
	for true {
		if index >= 5 {
			break
		}
		index++
		fmt.Printf("Why break index %d\n", index)
	}
}
