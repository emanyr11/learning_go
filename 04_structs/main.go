package main

import (
	"fmt"
	"strconv"
)

type person struct {
	fName, lName, gender, email, address string
	age, phNumber                        int
}

//Value Reciever
func (p person) greet() string {
	return "Hello i am " + p.fName + " " + p.lName + " and i am " + strconv.Itoa(p.age)
}

// Pointer Recievr * not returning modifying (no return type)
func (p *person) hasBirthday() {
	p.age++
}

// Pointer Recievr * not returning modifying passing in a spouselName string
func (p *person) getMarried(spouseLastname string) {
	if p.gender == "F" {
		p.lName = spouseLastname
	}
}

func main() {
	person1 := person{"Homer", "Simpson", "M", "HJSimpson@gmail.com", "742 Evergreen Terrace", 39, 5558628}
	person2 := person{"Marge", "Bouvier", "F", "MJBouvier@gmail.com", "742 Evergreen Terrace", 36, 5558628}
	//Prints our Struct
	fmt.Println(person1)
	fmt.Println(person2)
	//Print out greeting
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	// Homer ages 1 year
	person1.hasBirthday()
	// Marge Marries and becomes a Simpson
	person2.getMarried("Simpson")
	// Reprint Greeting with modifications
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
