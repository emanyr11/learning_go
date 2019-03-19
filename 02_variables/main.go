package main

import (
	"fmt"
)

func main() {
	name := "steve"
	age := 21
	canLegalyDrink := false
	if age >= 21 {
		canLegalyDrink = true
	}
	fmt.Println("Hi my name is: " + name)
	fmt.Println("I am: ", age)
	fmt.Println("Can i drink: ", canLegalyDrink)

}
