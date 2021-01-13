package main

import ("fmt"
		"strconv")

// Person - defnine Person struct
type Person struct{
//	firstName string
//	lastName string
//	city string 
//	gender string
//	age int

//altrnitively you can 
	firstName, lastName, city, gender string
	age 							   int

}

// methods. There are 2 types 1) vlue reciever and pointer reciever
// value reciever method are for methods that just do calculations and things (they wont change anything)
// pointer recievers are used when we want to change something.

// greeting method (value reviever)
func(p Person) greet() string{
	// note all these datatypes need to be the same so we will need to change age to a sting - note the strconv package
	return "Hello, My name is " + p.firstName + " " + p.lastName + " and I am a " + strconv.Itoa(p.age)
}
// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer Reciever)
func (p *Person) getMarried(spouseLastName string){
	if p.gender == "m"{
		return
	} else {
		p.lastName = spouseLastName // we will be passing spouseLastNamer in 
	}
}


func main() {
	//Init person using struct
	person1 := Person {firstName : "Samantha", lastName : "Bennette", city: "Dublin", gender: "f", age : 25}
	// alternitivly
	person2 := Person { "Bob", "Francisco",  "Miami",  "f", 47}

	//fmt.Println(person1, person2)
	//person1.age++
	//fmt.Println(person1.firstName, person1.age)

	// calling the greet method
	person1.hasBirthday()
	//person1.hasBirthday()
	//person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
	
	

}

