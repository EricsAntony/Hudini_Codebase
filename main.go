package main

import "fmt"

func main() {
	numbers := []float64{2, 3, 4, 5, 6, 7}
	fmt.Printf("Question 1:\nThe average is : %.1f\n\n", calculateAverage(numbers))
	checkAge(25)
	fmt.Printf("Question 3:\nThe string after reverse is %s", reverseString("hello"))
	fmt.Printf("\n\nQuestion 4:\nThe largerst number is %d\n\n", findLargestNumber([]int{10, 20, 30, 40, 50}))
	fmt.Printf("Question 5:\n")
	p := count{}
	p = p.increment()
	p = p.increment()
	p.display()
	p = p.decrement()
	p.display()
	p = p.decrement()
	p.display()
}

// // 1. Calculate Average
// Objective: Write a function that takes an array of numbers and returns the average. Use loops and basic arithmetic.
// calculateAverage calculates teh average of the slice and returns the average
func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	//loop through the slice and find sum;
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	// return the average
	return sum / float64(len(numbers))
}

// 2. Check Age
// Write a function that takes an age and prints whether the person is a minor, a young adult, or an adult.
// Use conditional statements.
// Function signature:
// checkAge check whether the age given is in a specific age range and prints valid response.
func checkAge(age int) {
	// Write your code here to determine and print if the age corresponds to a minor, a young adult, or an adult.
	if age < 18 {
		fmt.Printf("Question 2:\nThe Person is a Minor.\n\n")
	} else if age >= 18 && age < 25 {
		fmt.Printf("Question 2:\nThe Person is a Young adult.\n\n")
	} else {
		fmt.Printf("Question 2:\nThe Person is an Adult.\n\n")
	}
}

// 3. Reverse String
// Objective: Create a function that reverses a string. This will demonstrate basic string manipulation and for loops.
// Function signature:
// reverseString function reverses the given string and prints it.
func reverseString(str string) string {
	// Write your code here to reverse and return the string.
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// 4. Find Largest Number
// Objective: Write a function that takes an array of numbers and returns the largest number.
// Use loops and conditional statements to solve the problem.
// Function signature:
// findLargestNumber function takes array of numbers and returns the largest number
func findLargestNumber(numbers []int) int {
	// Write your code here to find and return the largest number in the array.
	largest := 0
	// Loop iterate over the slice and finds the largest number.
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}
	// returns the largest number.
	return largest
}

// 5. Simple Counter
// Create an object that acts as a counter with methods to add, subtract, and reset the count.
// Demonstrate object methods and the use of this.

type count struct {
	counter int
}

func (p count) increment() count {
	p.counter++
	return p
}

func (p count) decrement() count {
	p.counter--
	return p
}

func (p count) display() {
	fmt.Printf("\n\n%d ", p)
}


