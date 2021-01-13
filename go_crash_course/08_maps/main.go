package main

import "fmt"

/* is a collection of unordered pairs of key-value. 
It is widely used because it provides fast lookups and values that can retrieve, 
update or delete with the help of keys. It is a reference to a hash table
*/

func main() {
	// define map
	//map will have name for key emails for value
	//               [dt for key]dt for value)
	emails:= make(map[string]string)
	//assign key values

	emails["Bob"] = "bob@gmial.com"
	emails["Sharon"] = "sharon@gmial.com"
	emails["deleteMe"] = "deleteMe@gmial.com"

	//we can also declare a map
/*	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"Sharon@gmail.com"} 	*/


	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	///delete from map
	delete(emails, "deleteMe")
	fmt.Println(emails)

	//we can also declare a map
}