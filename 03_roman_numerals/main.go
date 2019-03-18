package main

import (
	"fmt"
)
import s "Strings"

func main() {
	rNumeral := "IV"
	rNumeralInt := 0
	fmt.Println("Enter roman numeral:")
	fmt.Scanf("%s", &rNumeral)
	fmt.Println("Numerical value of Roman Numeral:")

	// Using case switch to assign value based off the users input
	switch rNumeral {
	case "IV":
		fmt.Println("4")
	case "IX":
		fmt.Println("19")
	case "XXIX":
		fmt.Println("29")
	}
	// if IX -1
	// If IV -1
	// Using if else to see the contents of rNumeral and assigning value based on numeral count
	if (s.Count(rNumeral, "I")) == 1 {
		if rNumeral == "IV" {
			rNumeralInt += 4
		} else if rNumeral == "IX" {
			rNumeralInt += 9
		} else if (s.Count(rNumeral, "V")) == 1 {
			rNumeralInt += 5
			rNumeralInt++
		} else if (s.Count(rNumeral, "X")) == 1 {
			rNumeralInt += 10
			rNumeralInt++
		} else {
			rNumeralInt++
		}
	} else if (s.Count(rNumeral, "I")) == 2 {
		if (s.Count(rNumeral, "V")) == 1 {
			rNumeralInt += 5
			rNumeralInt += 2
		} else if (s.Count(rNumeral, "X")) == 1 {
			rNumeralInt += 10
			rNumeralInt += 2
		} else {
			rNumeralInt += 2
		}
	} else if (s.Count(rNumeral, "I")) == 3 {
		if (s.Count(rNumeral, "V")) == 1 {
			rNumeralInt += 5
			rNumeralInt += 3
		} else if (s.Count(rNumeral, "X")) == 1 {
			rNumeralInt += 10
			rNumeralInt += 3
		} else {
			rNumeralInt += 3
		}
	} else if (s.Count(rNumeral, "V")) == 1 {
		rNumeralInt += 5
	} else if (s.Count(rNumeral, "X")) == 1 {
		rNumeralInt += 10
	} else if (s.Count(rNumeral, "L")) == 1 {
		rNumeralInt += 50
	} else if (s.Count(rNumeral, "C")) == 1 {
		rNumeralInt += 100
	}

	fmt.Printf("Int = %d\n", rNumeralInt)
}
