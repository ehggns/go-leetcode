package main

import "fmt"

// dailyTemperatures returns an array where each element represents the number of days
// you would have to wait until a warmer temperature.
//
// For each day, it finds the next day with a higher temperature and returns the
// number of days to wait. If there is no future day with a higher temperature,
// the value remains 0.
//
// Example:
//
//	Input: temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
//	Output: [1, 1, 4, 2, 1, 1, 0, 0]
//
// The function uses a monotonic stack approach with O(n) time complexity and O(n) space complexity.
// Parameters:
//   - temperatures: A slice of integers representing daily temperatures.
//
// Returns:
//   - A slice of integers where each value represents how many days until a warmer temperature.
func dailyTemperatures(temperatures []int) []int {
	// Initialize a slice to store the result with the same length as temperatures
	result := make([]int, len(temperatures))
	// Initialize a stack to keep track of indices of temperatures
	stack := make([]int, 0)

	for i, temp := range temperatures {
		// While the stack is not empty and the current temperature is greater than the temperature at the index stored at the top of the stack
		print("outer i: ", i, " ", "  temp: ", temp, "\n")
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1] // Pop the index from the stack

			// debugging output
			print("j: ", j, "  i: ", i, " ", "  temperatures[j]: ", temperatures[j], " ", "  temp: ", temp, "\n")

			stack = stack[:len(stack)-1] // Remove the last index from the stack
			// Calculate the difference in days and store it in the result
			result[j] = i - j // The number of days until a warmer temperature
		}
		stack = append(stack, i) // Push the current index onto the stack
	}

	return result
}

func main() {
	// Example usage
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temperatures)
	for _, r := range result {
		fmt.Print(r) // Output: 1 1 4 2 1 1 0 0
	}
	fmt.Println()
	temperatures2 := []int{30, 40, 50, 60}
	result2 := dailyTemperatures(temperatures2)
	for _, r := range result2 {
		fmt.Print(r) // Output: 1 1 1 0
	}
	fmt.Println()
	temperatures3 := []int{51, 52, 56, 51, 53}
	result3 := dailyTemperatures(temperatures3)
	for _, r := range result3 {
		fmt.Print(r) // Output: 1 1 0 1 0
	}
	fmt.Println()
}
