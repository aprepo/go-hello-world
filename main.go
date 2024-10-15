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

	// This is a 'slice'
	var large_numbers []string
	var small_numbers []string
	for i = 0; i <= MAX_LOOP_COUNT; i++ {
		fmt.Print("Number " + strconv.Itoa(i))
		if i < 5 {
			small_numbers = append(small_numbers, strconv.Itoa(i)+" is a small number")
			fmt.Println(" is small number")
		} else {
			large_numbers = append(large_numbers, strconv.Itoa(i)+" is a small number")
			fmt.Println(" is large number")
		}
	}
	fmt.Println(`There were ` + strconv.Itoa(len(small_numbers)) + " small numbers")
	fmt.Println(`and ` + strconv.Itoa(len(large_numbers)) + " large numbers")

	// Iterate over a slice. The slice operator looks very Python
	for j, number := range small_numbers[1:4] {
		fmt.Print(j)
		fmt.Println(number)
	}
}
