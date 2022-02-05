package main

import (
	"fmt"
)

//global
var number int

func varibles() {

	number = 124

	//only plus
	var plusnumber uint
	plusnumber = 32323

	//double
	var doublenumber float32

	doublenumber = 377777777777777777.323

	fmt.Println(number)
	fmt.Println(plusnumber)
	fmt.Println(doublenumber)

	a, b, c := 1, 2, 3

	a, b = 4, 5

	a, b = b, a

	a, _, b = 1, 2, 3

	fmt.Println(a, b, c)
}
