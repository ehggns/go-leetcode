package main

// findMinimun finds the minimum element in a rotated sorted array.
//
// A rotated sorted array is an array that was originally sorted in non-decreasing order (each element is greater than or equal to the previous element), but then was "rotated" by shifting all elements to the right by some number of positions, with the elements that fall off the end wrapping around to the beginning.
//
// For example, [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2] shifting 4 positions to the right.
//
// Characteristics of a Rotated Sorted Array:
//
//  1. Original Order: The array was initially sorted in ascending order.
//
//  2. Rotation Process: The array was rotated by shifting elements k positions to the right (or equivalently, n-k positions to the left), where k can be any integer between 1 and n-1, and n is the length of the array.
//
//  3. Structure: After rotation, the array consists of two sorted subarrays, with the elements of the second subarray being smaller than the elements of the first subarray.
//
// The function uses binary search to find the minimum value with O(log n) time complexity.
//
// Parameters:
//   - nums: A non-empty rotated sorted array of integers with no duplicates
//
// Returns:
//   - The minimum element in the array
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
func findMinimun(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	// Example usage
	nums := []int{3, 4, 5, 1, 2}
	result := findMinimun(nums)
	println("Minimum element:", result) // Output: Minimum element: 1
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	result2 := findMinimun(nums2)
	println("Minimum element:", result2) // Output: Minimum element: 0
	nums3 := []int{4, 5, 6, 7, 9, 0, 1, 2, 3}
	result3 := findMinimun(nums3)
	println("Minimum element:", result3) // Output: Minimum element: 0
}
