package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TommyFiga/go-mathy/calc"
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

	num1, err := calc.ConvertToNumber(input1)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := calc.ConvertToNumber(input2)
	if err != nil {
		log.Fatal(err)
	}

	res, err := calc.Calculator(num1, operator, num2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %g\n", res)
}