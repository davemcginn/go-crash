package main

import "fmt"

// function name(paramater type) return type
func greeting(name string) string{
	return "Hello " + name
}

//func getSub(num1 int, num2 int) int{ - or, as they are both int's
	func getSub(num1, num2 int) int{
	return num1 + num2

}

func main() {
	fmt.Println(greeting("Dave"))
	fmt.Println(getSub(5,5))
}