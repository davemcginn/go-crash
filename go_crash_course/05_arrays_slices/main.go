package main

import "fmt"
//arrays are fixed size, slices are not

//Arrays (in Go, they need 1: to be names and 2:to be of a fixed lenght.
//being a fixed lenght can be an issue so we can use slices.
func main() {
//declare then assing value
/*	var fruitArr [2]string
*/
//Assign values
/*	fruitArr[0] = "Orange"
	fruitArr[1] = "Apple"
*/
//declare and assing
/*	tropicFruitArr := [2]string{"mango", "pineapple"}
*/

/*	fmt.Println(fruitArr, tropicFruitArr)
*/
//Slices
vegSlice := []string{"potato", "carrot","Broccli"}


fmt.Println(vegSlice)
fmt.Println(len(vegSlice)) //count the number of elements
fmt.Println(vegSlice[1:2]) // range will not print out the last element.

}