package main

import "fmt"

func main() {
	x := 10
	y := 10
	//if (x < y) using perentises will run but it is commmon practice in go not to use them.
	
	//if else
	if x <= y{
		fmt.Printf("%d is less then or equal to %d \n",x , y)
	} else{
		fmt.Printf("%d is less then %d \n",y , x)
	}

	//else if
	color := "green"
		if color == "red"{
			fmt.Println("color is red")
		} else if color == "blue"{
			fmt.Println("color is blue")
		} else  {
			fmt.Println("color isnt blue or red")
		}

	//switch
	switch color{
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")

	}
}