/*
containsDuplicate determines if the input slice contains any duplicate elements.
It uses a map as a set data structure to track seen elements for O(1) lookups.

Parameters:
  - nums: A slice of integers to check for duplicates

Returns:
  - true if any value appears at least twice in the slice
  - false if every element is distinct

Time Complexity: O(n) where n is the length of nums
Space Complexity: O(n) in the worst case when all elements are distinct
*/
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
