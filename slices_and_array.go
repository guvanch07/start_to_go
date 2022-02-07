package main

import (
	"fmt"
)

func main() {
	// array
	// arrayEmpty := [3]string{"1", "2", "3"}

	// fmt.Println(arrayEmpty)
	// //slice
	// slicearray := []string{}
	// fmt.Println(slicearray)

	// message := make([]string, 5)
	// fmt.Println(len(message))
	// fmt.Println(cap(message))
	// message = append(message, "6")
	// fmt.Println(len(message))
	// fmt.Println(cap(message))
	doubleSlice()
	matrixSlice()
	matrixEseaVersion()
}

//двойной слайс

func doubleSlice() {
	matrix := make([][]int, 10)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x
		}
		fmt.Println(matrix[x])
	}
}

func matrixSlice() {
	matrix := make([][]int, 10)

	counter := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	}
}

func matrixEseaVersion() {
	mess := []string{
		"res1", "res2", "res3", "res4", "res5", "res6",
	}

	for i := range mess {
		fmt.Println(mess[i])
	}
}
