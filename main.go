package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	log.SetPrefix("go-mathy: ")
	log.SetFlags(0)

	args := os.Args

	if len(args) != 4 {
		log.Fatal("invalid number of arguments")
	}

	input1 := args[1]
	operator := args[2]
	input2 := args[3]

	num1, err := convertToNumber(input1)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := convertToNumber(input2)
	if err != nil {
		log.Fatal(err)
	}

	res, err := Calculator(num1, operator, num2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %g\n", res)
}