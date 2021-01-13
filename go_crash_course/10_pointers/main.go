package main

import "fmt"

//pointers give you the ability toi point to the memory address/ location of a value of a variable

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	// get datatype
	fmt.Printf("%T\n", b) // this will produce *int the * represents a pinter

	// use to read value form address
	fmt.Println(*b)
	fmt.Println(*&a) //in eesnce this is the same


	//We can also change the value with pointer

	*b = 10
	fmt.Println(a)
}