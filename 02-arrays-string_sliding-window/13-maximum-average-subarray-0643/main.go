package main

// findMaxAverageSubarray finds the maximum average value of a contiguous subarray of length k in the given array.
//
// This function implements a sliding window approach to find the maximum average:
// 1. It first validates input parameters, returning 0 for invalid inputs
// 2. It calculates the sum of the first k elements
// 3. Then slides the window through the array, updating the current sum and tracking the maximum
// 4. Finally returns the maximum average by dividing the maximum sum by k
//
// Parameters:
//   - nums: The input integer array
//   - k: The size of the subarray to consider
//
// Returns:
//   - float64: The maximum average value of a contiguous subarray of length k
//
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(1) as it uses only constant extra space
func findMaxAverageSubarray(nums []int, k int) float64 {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return 0
	}

	maxSum, currentSum := 0, 0

	// Sliding window to calculate the sum of the first 'k' elements
	start, end := 0, k

	for end < len(nums) {
		// Remove previous element from the sum
		currentSum -= nums[start]
		start++

		// Add next element to the sum
		currentSum += nums[end]
		end++

		// Update maxSum if currentSum is greater
		maxSum = max(maxSum, currentSum)
	}
	return float64(maxSum) / float64(k)
}


func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4 // Size of the subarray
	result := findMaxAverageSubarray(nums, k)
	println(result) // Output: 12.75
}