package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	greetings := hello("dave")
	fmt.Println(greetings)
}

func hello(s string) string {
	greet := fmt.Sprintf("Hello ")
	return greet + s
}
