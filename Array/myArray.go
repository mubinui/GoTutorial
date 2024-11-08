package main

import "fmt"

func main(){

	fmt.Println("Welcome to the arrays in GoLang ")

	var fruitList [4] string  // in go we need to explicitly say what should be the array size \

	fruitList[0] = "Apple"
    fruitList[2] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruitlist is : ",fruitList)
	//Fruitlist is :  [Apple  Orange Banana]
	// the most weird thing about the go lang ...but we assigned the memory so length is actually
	fmt.Println("Lets print the length of the array ", len(fruitList))
	var vegList = [3] string{"potato","beans", "eggplant"}
	fmt.Println(vegList)

	// there isn't sorting or other functionality in Go-Lang 


}