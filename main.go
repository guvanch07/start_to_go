package main

import (
	"errors"
	"fmt"
)

func main() {

	mess := "Soon i become go"
	fmt.Println(mess)
	changeMessage(&mess)
	fmt.Println(mess)

	// //замкание

	// inc := increment()

	// fmt.Println(inc())
	// fmt.Println(inc())
	// fmt.Println(inc())
	// fmt.Println(inc())

	// fmt.Println(findmin(23, 3232, 32323, 4545, -434, -23, 23))

	// var message string
	// message = "i started to learn go"
	// message = "good luck"
	// fmt.Println(message)

	//varibles()
	// _functions("function is ready")
	// somethingfunc := returnstring("guvanch", 23)
	// println(somethingfunc)

	// age, whatis := enterToThe(34)

	// //age, _ := enterToThe(34)
	// println(age, whatis)

	// howold, err := enterToTheWithError(12)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(howold)

	// // ананимные функции
	// func() {
	// 	fmt.Println("unnoumous func")
	// }()

}

////functions

func _functions(varible string) {

	fmt.Println(varible)
}

// return var
func returnstring(name string, age int) string {
	concatanation := fmt.Sprintf("hello, %s! you %d", name, age)
	return concatanation
}

// return two or more varibles
func enterToThe(age int) (string, bool) {
	if age > 18 {
		return "come in but be ceerfully", true
	} else {
		return "stop", false
	}

}

//handler error from pck error
func enterToTheWithError(age int) (string, error) {
	if age > 18 {
		return "come in but be ceerfully", nil
	} else {
		return "stop", errors.New("smth is wrong")
	}

}

// find the smaller number

func findmin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

// замкание

func increment() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// указатели *deref &ref

func changeMessage(mess *string) {
	*mess += " from func changemessage"
}
