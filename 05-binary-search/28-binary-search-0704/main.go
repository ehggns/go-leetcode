package main

// binarySearch performs a binary search on a sorted array to find the target value.
// It returns the index of the target if found, or -1 if not present in the array.
//
// The function uses a two-pointer approach to narrow down the search space.
// It calculates the middle index using the formula:
//
//   - mid = left + (right-left)/2
//
// This formula is used to prevent overflow in (left + right) that can occur with the simpler formula:
//
//   - mid = (left + right)/2
//
// It is a mathematical equivalence because:
//
// left + (right-left)/2
//
// = left + right/2 - left/2
//
// = left/2 + right/2
//
// = (left + right)/2
//
// Parameters:
//   - nums: A sorted integer array to search in
//   - target: The integer value to search for
//
// Returns:
//   - The index of the target in the array if found, otherwise -1
//
// Time Complexity: O(log n) where n is the length of nums
// Space Complexity: O(1)
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // Avoids potential overflow
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	// Example usage
	nums := []int{-1, 0, 3, 5, 9, 12} // Sorted array
	target := 9
	result := binarySearch(nums, target)
	println("index:", result, "target:", nums[result]) // Output: index: 4, target: 9 (index of target in nums)
}
