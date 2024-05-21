package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {

	// TODO: Implement Calculator function
	// Checking whether the operand is valid. If not valid, return an error.
	if operand != "+" && operand != "-" && operand != "*" && operand != "/" {
		return 0.0, errors.New("Invalid arithmetic operation")
	}
	// Doing mathematical operations
	switch operand {
	case "+":
		return (value1 + value2), nil // Addition
	case "-":
		return (value1 - value2), nil // Subtraction
	case "*":
		return (value1 * value2), nil // Product
	case "/":if value2 == 0 {
		return 0.0, errors.New("Division by zero is not possible");
	}
		return (value1 / value2), nil // Division
	default: return 0.0, errors.New("Something went wrong") 
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	
	// TODO: Check if exact 3 parts are given. If not, throw an error
	if len(parts) != 3 {
		fmt.Print("No required input values")
		return
	}


	// TODO: Take all 3 values and pass to calculator function.
	// Convert the string type of parts to float.
	operand1, convertErr := strconv.ParseFloat(parts[0], 64)
	operand2, convertErr1:= strconv.ParseFloat(parts[2], 64)

	// Checking whether any of the inputs are not a number. 
	if convertErr != nil || convertErr1 != nil {
		fmt.Printf("Error: Not a number");
		return;
	}

	// Calling the Calculator function and stores the result in result variable and error in error variable.
	result, err := Calculator(operand1, operand2, parts[1])

	// Handling error from the Calculator function.
	if err != nil {
		fmt.Printf("Error: %v", err);
		return;
	}

	// TODO: Print results
	fmt.Printf("Result: %v\n", result)
}
