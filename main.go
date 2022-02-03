package main

import (
	"fmt"
)

func main() {
	var message string
	message = "i started to learn go"
	message = "good luck"
	fmt.Println(message)

	_varibles()
	_functions("function is ready")
	somethingfunc := returnstring("guvanch", 23)
	println(somethingfunc)

	age, whattodo := enterToThe(34)
	//age, _ := enterToThe(34)
	println(age, whattodo)

}

//global
var number int

func _varibles() {

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

////functions

func _functions(varible string) {

	fmt.Println(varible)
}

func returnstring(name string, age int) string {
	concatanation := fmt.Sprintf("hello, %s! you %d", name, age)
	return concatanation
}

func enterToThe(age int) (string, bool) {
	if age > 18 {
		return "come in but be ceerfully", true
	} else {
		return "stop", false
	}

}
