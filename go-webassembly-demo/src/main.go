package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	fmt.Println("Go WebAssembly Demo")

	// Register a function to be called from JavaScript
	js.Global().Set("addFunction", js.FuncOf(addFunction))

	// Keep the program running indefinitely
	select {}
}

func addFunction(this js.Value, args []js.Value) interface{} {
	// Get the input value from JavaScript
	inputString1 := args[0].String()
	inputString2 := args[1].String()

	// Perform some computation
	result := compute(inputString1, inputString2)

	// Return the result to JavaScript
	return js.ValueOf(result)
}

func compute(inputString1 string, inputString2 string) string {

	// Convert input strings to integers
	num1, err := strconv.Atoi(inputString1)
	if err != nil {
		return "Invalid inputString1 :" + inputString1
	}

	num2, err := strconv.Atoi(inputString2)
	if err != nil {
		return "Invalid inputString2 :" + inputString2
	}

	// Perform addition
	addition := num1 + num2
	// Log the addition result
	fmt.Println("Addition result:", addition)
	return strconv.Itoa(addition)
}
