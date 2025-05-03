package main

// maxArea calculates the maximum area of water that can be contained between two vertical lines
// in a container represented by an array of heights.
//
// The function uses the two-pointer technique to find the maximum area:
// - It starts with two pointers at the leftmost and rightmost positions
// - For each position, it calculates the area between the lines
// - It moves the pointer with the smaller height inward
// - This process continues until the pointers meet
// - The maximum area found during this process is returned.
//
// Time Complexity: O(n) - where n is the number of elements in the heights array
// Space Complexity: O(1) - no additional space is used that scales with input size
// Example:
// Given heights = [1,8,6,2,5,4,8,3,7]
// The maximum area is 49, formed between the lines at indices 1 and 8 (height 8 and 7).
// The area is calculated as min(8, 7) * (8 - 1) = 7 * 7 = 49.
// Parameters:
//   - heights: []int - slice of integers representing heights of vertical lines
//
// Returns:
//   - int - the maximum area of water that can be contained
func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	maxArea := 0

	for left < right {
		width := right - left
		height := min(heights[left], heights[right]) // Find the minimum height between the two pointers
		area := width * height
		maxArea = max(maxArea, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(heights)
	println(result) // Output: 49
}
