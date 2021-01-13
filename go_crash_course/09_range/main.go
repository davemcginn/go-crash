package main

import "fmt"

//range is used to loop through arrays, mpas, slices...

func main() {
	ids := []int{33,56,77,88,65,43,2}

	// loop through id's
	for i, id := range ids {

		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	// loop through id's
	for _, id := range ids {  // if you use  "for i..."" you are declaring a var i that you have to use. Use an "_" to get atounfd this.

		fmt.Printf("ID: %d\n", id)
	}


	// add id's together
	sum := 0 
	for _, id := range ids{
		sum += id
	}

	//range with map
	email := map[string]string{ "John":"john:gmail.com", "Mary" : "mary@gmail.com"}

//	fmt.Println(email)
	for k, v := range email{
		fmt.Printf("%s: %s\n", k, v)
	}
//you can do this with just the key aswell.
	for k := range email{
		fmt.Println("my name is", k)
	}
}