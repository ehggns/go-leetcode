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

// ## Detailed Walkthrough of `{73, 74, 75, 71, 69, 72, 76, 73}`

// ### Starting Point:
// - Result array (initialized): `[0, 0, 0, 0, 0, 0, 0, 0]`
// - Stack (empty): `[]`

// ### Day 0 (temp = 73):
// - Stack is empty, so no inner loop execution
// - Push index 0 to stack
// - Stack: `[0]`

// ### Day 1 (temp = 74):
// - Current temp (74) > temp at stack top (73)
// - Enter inner loop:
//   - Pop index 0 from stack
//   - Set result[0] = 1 - 0 = 1 (wait 1 day for warmer temp)
//   - Stack becomes empty, exit inner loop
// - Push index 1 to stack
// - Result: `[1, 0, 0, 0, 0, 0, 0, 0]`
// - Stack: `[1]`

// ### Day 2 (temp = 75):
// - Current temp (75) > temp at stack top (74)
// - Enter inner loop:
//   - Pop index 1 from stack
//   - Set result[1] = 2 - 1 = 1 (wait 1 day for warmer temp)
//   - Stack becomes empty, exit inner loop
// - Push index 2 to stack
// - Result: `[1, 1, 0, 0, 0, 0, 0, 0]`
// - Stack: `[2]`

// ### Day 3 (temp = 71):
// - Current temp (71) < temp at stack top (75)
// - Skip inner loop
// - Push index 3 to stack
// - Result: `[1, 1, 0, 0, 0, 0, 0, 0]`
// - Stack: `[2, 3]`

// ### Day 4 (temp = 69):
// - Current temp (69) < temp at stack top (71)
// - Skip inner loop
// - Push index 4 to stack
// - Result: `[1, 1, 0, 0, 0, 0, 0, 0]`
// - Stack: `[2, 3, 4]`

// ### Day 5 (temp = 72):
// - Current temp (72) > temp at stack top (69)
// - Enter inner loop:
//   - Pop index 4 from stack
//   - Set result[4] = 5 - 4 = 1 (wait 1 day for warmer temp)
//   - Stack: `[2, 3]`
//   - Check next: 72 > 71 at stack top (index 3)
//   - Pop index 3 from stack
//   - Set result[3] = 5 - 3 = 2 (wait 2 days for warmer temp)
//   - Stack: `[2]`
//   - 72 < 75 at stack top (index 2), exit inner loop
// - Push index 5 to stack
// - Result: `[1, 1, 0, 2, 1, 0, 0, 0]`
// - Stack: `[2, 5]`

// ### Day 6 (temp = 76):
// - Current temp (76) > temp at stack top (72)
// - Enter inner loop:
//   - Pop index 5 from stack
//   - Set result[5] = 6 - 5 = 1 (wait 1 day for warmer temp)
//   - Stack: `[2]`
//   - Check next: 76 > 75 at stack top (index 2)
//   - Pop index 2 from stack
//   - Set result[2] = 6 - 2 = 4 (wait 4 days for warmer temp)
//   - Stack becomes empty, exit inner loop
// - Push index 6 to stack
// - Result: `[1, 1, 4, 2, 1, 1, 0, 0]`
// - Stack: `[6]`

// ### Day 7 (temp = 73):
// - Current temp (73) < temp at stack top (76)
// - Skip inner loop
// - Push index 7 to stack
// - Result: `[1, 1, 4, 2, 1, 1, 0, 0]`
// - Stack: `[6, 7]`

// ### Final Result: `[1, 1, 4, 2, 1, 1, 0, 0]`

// This means:
// - Day 0 (73°): Wait 1 day for a warmer temperature (74° on day 1)
// - Day 1 (74°): Wait 1 day for a warmer temperature (75° on day 2)
// - Day 2 (75°): Wait 4 days for a warmer temperature (76° on day 6)
// - Day 3 (71°): Wait 2 days for a warmer temperature (72° on day 5)
// - Day 4 (69°): Wait 1 day for a warmer temperature (72° on day 5)
// - Day 5 (72°): Wait 1 day for a warmer temperature (76° on day 6)
// - Day 6 (76°): No warmer temperatures ahead (remains 0)
// - Day 7 (73°): No warmer temperatures ahead (remains 0)
