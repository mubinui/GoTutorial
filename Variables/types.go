package main

import "fmt"

const LoginToken string = "dffgerferftdrfdtfdrterterrgfdr"

func main (){
	var username string = "mubin"
	fmt.Println("username")
	fmt.Printf("The type of the variable is: %T \n", username)

	// there is a difference between println and printf

	var x = 10;
	fmt.Printf("%T \n",x)

	//boolean 

	var isLoggedIn = false
	fmt.Printf("%T \n",isLoggedIn)

	// uint8

	var smallVal uint8 = 156
	// var smallVal2 uint8 = 277 not possible
	fmt.Printf("%T \n", smallVal)

	//float
	var smallFloat float32 = 255.45353463664366436
	fmt.Println(smallFloat)
	fmt.Printf("%8f \n", smallFloat)

	// default values and some aliases

	var another int
	fmt.Println(another)
	// the outout of this will be 0

	// no var
	users :=1000
	fmt.Println(users)
	fmt.Printf("%T \n", users)
}