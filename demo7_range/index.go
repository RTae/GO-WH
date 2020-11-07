package main

import (
	"fmt"
)

func main() {
	course := []string{"Android", "IOS", "React"}

	for index, item := range course {
		fmt.Printf("%d, %s\n", index+1, item)
	}

	for _, item := range course {
		fmt.Printf("%s\n", item)
	}
}
