package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Begin")

	// Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true
	tmp1 = 10
	const con1 int = 2
	//con1 = 10

	//Implicit Declaration
	// var tmp5 int = 0
	tmp5 := 0
	tmp6 := true
	tmp7 := 24.65
	tmp8 := "Test implicit"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(con1)

	fmt.Println(tmp5)
	fmt.Println(tmp6)
	fmt.Println(tmp7)
	fmt.Println(tmp8)

	run()
}

func run() {
	count++
	fmt.Println(count)
}
