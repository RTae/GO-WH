package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["two"])

	var color = make(map[string]string)
	color["red"] = "#f00"
	color["gree"] = "#0f0"
	color["blue"] = "#00f"
	fmt.Println("", color)
	fmt.Println("", color["green"])

	var course = make(map[string]map[string]int)

	course["Android"] = make(map[string]int)
	course["Android"]["price"] = 200
	course["Android"]["code"] = 1

	course["IOS"] = make(map[string]int)
	course["IOS"]["price"] = 300
	course["IOS"]["code"] = 2

	fmt.Println(course)
	fmt.Println(course["Android"])
	fmt.Println(course["IOS"])
}
