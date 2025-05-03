package main

/*
twoSum finds two indices in the given array whose values sum up to the target.
It uses a hash map to store the complement of each number and its index.
The function returns a slice of two integers representing the indices of the numbers that sum to target.
If no solution is found, it returns nil.

Parameters:
  - nums: A slice of integers to search through
  - target: The target sum to find

Returns:
  - []int: A slice containing two indices whose corresponding values sum to target,
    or nil if no solution exists

Time complexity: O(n) where n is the length of nums
Space complexity: O(n) for the hash map storage
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)     // Create a map to store the indices of the numbers
	for i, num := range nums { // Iterate through the array
		complement := target - num // Calculate the complement
		if j, exists := m[complement]; exists {
			return []int{j, i} // Return the indices of the two numbers
		}
		m[num] = i // Store the index of the current number
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15} // Example input
	target := 9                 // Example target
	result := twoSum(nums, target)
	println(result) // Output: [0 1]
}
