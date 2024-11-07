package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcom to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for mly skill : ")

	//comma ok || err ok 
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks of the rating ", input)

	newReader := bufio.NewReader(os.Stdin)

	input2,err := newReader.ReadString('\n')

	numrating ,err2 := strconv.ParseInt(strings.TrimSpace(input2),10,32) 
	
	if err2 != nil {
		fmt.Println(err2)
	}else{
		fmt.Println(numrating+100)
	}

	fmt.Println(err)

}