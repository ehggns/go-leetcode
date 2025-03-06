package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}) // Create a set to store the numbers (using a map and struct{} to save space)
	for _, num := range nums {  // Iterate through the array
		if _, exists := m[num]; exists { // Check if the number already exists in the set
			return true // If it exists, return true
		}
		m[num] = struct{}{} // Otherwise, add the number to the map (set initialization)
	}
	return false // If no duplicates were found, return false
}

func main() {
	nums := []int{1, 2, 3, 1} // Example input
	result := containsDuplicate(nums)
	fmt.Println(result) // Output: true

	nums = []int{1, 2, 3, 4} // Example input
	result = containsDuplicate(nums)
	fmt.Println(result) // Output: false
}
