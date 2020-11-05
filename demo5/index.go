package main

import (
	"fmt"
)

func main() {
	fn1()
	fn2("Good morning", "1")
	const a, b = 2, 3
	x, y := swap(a, b)
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))
	fmt.Printf("%d,%d => %d,%d\n", a, b, x, y)

	_x, _y, title := swapv2(10, 20)
	fmt.Printf("%d,%d => %d,%d,%s", 10, 20, _x, _y, title)
}

func fn1() {
	fmt.Println("Codemobiles")
}

func fn2(title string, version string) {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapv2(a, b int) (x, y int, title string) {
	x = b
	y = a
	title = "result"
	return
}
