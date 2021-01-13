package main

import "fmt"

func main() {
	// long method

	i := 1
	for i <= 10{
		fmt.Println(i)
		//i = i + 1
		i++
	}

	//short method
	for i := 1; i <= 10; i++{
		fmt.Printf("Number %d\n", i )
	}

	//FizzBuzz
	for i:=1; i<= 100; i++{
		if i % 15 == 0{
			fmt.Printf("fizzBuzz Number %d\n", i )
		}else if i % 3 == 0{
			fmt.Printf("Fizz Number %d\n", i )
		}else if i % 15 == 0{
			fmt.Printf("Buzz Number %d\n", i )
		}else {
			fmt.Printf("Number %d\n", i )
		}
	}
}