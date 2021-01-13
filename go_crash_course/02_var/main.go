package main

import "fmt"

//shoes := "nike" // please see shorthand below err: non-declaration statement outside function bodygo 

func main() {
	// see datatypes here https://www.tutorialspoint.com/go/go_data_types.htm
	// In go when you create a vartiable you have to use it. Otherwiaw you will get "var declared but not used"

	//--**Creating a variable using var
	// no need to declare datatype as it is inferred var name string = "David McGinn"
	var name = "David McGinn"
	var age int32 = 30 // again int is not needed here as the datatype is inferred but we can specifically cast it as we pease
	//const
	var isCool = true //just leaving this as var so i dont throw an error
	isCool = false //if this was a var I can change it after the fact - you cant do this with a const
	//shorthand
	shoes := "nike" //you can only use this inside a finction
	size := 1.5
	// shorter hand
	shirt, nextSize := "Hugo Boss", 17.5

	fmt.Println(name, age, isCool, shoes, size, shirt, nextSize)

	// To get the dataype we can use a verb provided by fmt - https://golang.org/pkg/fmt/	
	fmt.Printf("%T\n", size) 
}

