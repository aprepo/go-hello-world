package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello world")

	fmt.Println("Try out different variables and stuff")
	var a = "a"
	fmt.Println("a: " + a)

	var b bool = true
	fmt.Println("b: " + strconv.FormatBool(b))

	var i int = 1
	fmt.Println("i: " + strconv.Itoa(i))

	j := 2 // No var needed if using :=
	fmt.Println("j: " + strconv.Itoa(j))

	const c int = 10
	fmt.Println("c: " + strconv.Itoa(c) + " is a constant and can't be modified")

	const MAX_LOOP_COUNT = 10
	for i = 0; i <= MAX_LOOP_COUNT; i++ {
		fmt.Print("Number " + strconv.Itoa(i))
		if i < 5 {
			fmt.Println(" is small number")
		} else {
			fmt.Println(" is large number")
		}
	}
}
