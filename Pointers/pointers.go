package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers") 
	// When we pass a value via a variable then it passes a copy of it if it changes copy doesn't change at the same time 
	// So sending the memory address can be more convanient ...and it will pass the actual value 
	// * is the symbol of pointer and next to the datatype says what kind of data it will hold 


	// var ptr *int

	// fmt.Println("Value of pointer is ",ptr)
	// fmt.Printf("%T \n",ptr)

	myNumber := 23

	var ptr2 = &myNumber // & symbol is the give the address of the memory address of the value 

	fmt.Println("Printing my number",*ptr2) // *ref meaning what is in the pointer 

	*ptr2 = *ptr2 + 2

	fmt.Println("The new value of myNumber is ",myNumber)




}