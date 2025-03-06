package main

import (
	"fmt"
)

// twoSum finds two numbers in nums that add up to target and returns their indices.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // Create a map to store the indices of the numbers
	for i, num := range nums { // Iterate through the array
		if j, ok := m[target-num]; ok { // Check if the complement exists
			return []int{j, i} // Return the indices of the two numbers
		}
		m[num] = i// Store the index of the current number
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15} // Example input
	target := 9 // Example target
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0 1]
}