package main

import (
	"fmt"
	"os"
	"strconv"

	"funcname/arithmetic" // using Div function arithmetic package of funcname workspace
)

func divide(a, b int) float32 {
	if div, err := arithmetic.Div(a, b); err != nil {
		// if any error, panic the error
		panic(err)
	} else {
		return div
	}
}

func getInput() (int, int) {
	if len(os.Args) < 3 {
		// panic can accept type of any value
		panic("minimum 3 arguments required")
	}

	num1, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		panic(err)
	}

	return int(num1), int(num2)
}

func main() {
	defer func() {
		switch err := recover().(type) {
		case string:
			fmt.Println("ERROR", err)
			os.Exit(1)
		case error:
			fmt.Println("ERROR", err.Error())
			os.Exit(1)
		}
		// // or you can simply print
		// if err := recover(); err != nil {
		// 	fmt.Println("ERROR:", err)
		// 	os.Exit(1)
		// }
	}()

	a, b := getInput()
	fmt.Println(a, "/", b, "->", divide(a, b))
}
