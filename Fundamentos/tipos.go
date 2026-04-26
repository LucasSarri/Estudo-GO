package main

import "fmt"

func main () {
	//bool
	fmt.Printf("Type: %T - Value: %v \n", true, true)
	fmt.Printf("Type: %T - Value: %v \n", false, false)
	//string
	fmt.Printf("Type: %T - Value: %v \n", "Lucas", "Lucas")
	fmt.Printf("Type: %T - Value: %v \n", "2", "2")
	//int
	fmt.Printf("Type: %T - Value: %v \n", 1, 1)
	fmt.Printf("Type: %T - Value: %v \n", 12354, 12354)
	//float
	fmt.Printf("Type: %T - Value: %v \n", 1.2333, 1.2333)
	fmt.Printf("Type: %T - Value: %v \n", 0.2, 0.2)
	//Printf
	//o %T apresenta o tipo de um valor e o %v apresenta o valor em si, independente do tipo
}