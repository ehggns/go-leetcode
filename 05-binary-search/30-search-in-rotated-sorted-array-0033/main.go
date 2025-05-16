package main

import "fmt"

func searchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] { // When left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Target is in the left half
			} else {
				left = mid + 1 // Target is in the right half
			}
		} else { // When right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Target is in the right half
			} else {
				right = mid - 1 // Target is in the left half
			}
		}
	}
	return -1 // Target not found
}

func main() {
	// Example usage
	nums := []int{4, 5, 6, 7, 9, 0, 1, 2, 3}
	target := 0
	result := searchInRotatedSortedArray(nums, target)
	fmt.Println("target is in position:", result) // Output: 5
}
