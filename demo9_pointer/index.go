package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)

	changeMessage(&msg, "Hi")
	fmt.Println(msg)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
