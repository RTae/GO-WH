package main

import "demo17/employee"

func main() {
	e := employee.Init("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()

	e = employee.Init("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
