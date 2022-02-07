package main

import (
	"fmt"
)

func main() {

	defer printmess()
	//fmt.Println("main()")

	mess := []string{
		"res1", "res2", "res3", "res4", "res5", "res6",
	}

	mess[6] = "res7"
	fmt.Println(mess)

}
func printmess() {
	if r := recover(); r != nil {
		fmt.Println()
	}
	fmt.Println("exseption!!!!")
}
