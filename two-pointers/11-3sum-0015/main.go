package main

import "sort"

// threeSum finds all unique triplets in the input array that sum up to zero.
// It uses a two-pointer technique after sorting the array to efficiently find triplets.
//
// Parameters:
//   - nums: An integer slice containing the input numbers
//
// Returns:
//   - [][]int: A 2D slice containing all unique triplets that sum to zero.
//     Each inner slice contains exactly three integers.
//     Returns nil if input is nil or has less than 3 elements.
//
// Time Complexity: O(n²) where n is the length of nums
// Space Complexity: O(1) excluding the space used for output
//
// Note: Go does not have buillt-in set data structure.
//       We use a map to track unique triplets.
//       The triplet is stored as a fixed-size array of 3 integers,
// 			 and the map's key is the triplet itself.
//       This ensures that we only add unique triplets to the result.
//       The map's value is an empty struct, which is a common Go idiom
//       for representing a set without needing to allocate memory for the value.
//       The map's key is a fixed-size array of 3 integers, which is a valid
//       key type in Go. This allows us to use the triplet as a key in the map
//       to check for uniqueness efficiently.
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	result := make([][]int, 0)
	seen := make(map[[3]int]struct{})
	for i := range len(nums) - 2 {
		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				triplet := [3]int{nums[i], nums[left], nums[right]}
				// Check if the triplet is already seen
				if _, ok := seen[triplet]; ok {
					left++
					right--
					continue
				}
				seen[triplet] = struct{}{}
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
			if target < 0 {
				left++
			}
			if target > 0 {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	for _, triplet := range result {
		println(triplet[0], triplet[1], triplet[2])
	}
}
