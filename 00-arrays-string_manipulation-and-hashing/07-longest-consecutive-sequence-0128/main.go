/*
longestConsecutiveSequence finds the length of the longest consecutive elements sequence
in an unsorted array of integers.

The function uses a map-based approach to achieve O(n) time complexity
(i.e. hashing technique).
1. All numbers are stored in a map for O(1) lookup
2. For each number, it explores consecutive sequences in both directions (forward and backward)
3. Numbers are marked with boolean values as visited to avoid redundant processing

Parameters:
  - nums: A slice of integers (can be unsorted and contain duplicates)

Returns:
  - The length of the longest consecutive sequence

Example:

	Input: [100, 4, 200, 1, 3, 2]
	Output: 4 (the consecutive sequence is [1, 2, 3, 4])

Time Complexity: O(n) where n is the length of the input array
Space Complexity: O(n) for the hash map
*/
package main

func longestConsecutiveSequence(nums []int) int {
	numsMap := make(map[int]bool)

	for _, num := range nums {
		numsMap[num] = false // mark as not visited
	}

	longestStreak := 0

	for _, num := range nums {
		if numsMap[num] {
			continue // already visited, skip this number
		}

		numsMap[num] = true // mark as visited
		currentNum := num
		currentStreak := 1
		// Check for consecutive numbers greater than the current number
		for numsMap[currentNum+1] { // loop while exists a consecutive number forwards
			currentNum++
			currentStreak++
			numsMap[currentNum] = true // mark as visited
		}

		// Check for consecutive numbers less than the current number
		currentNum = num
		for numsMap[currentNum-1] { // loop while exists a consecutive number backwards
			currentNum--
			currentStreak++
			numsMap[currentNum] = true // mark as visited
		}

		// Update the longest streak if needed
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}
	// Return the longest streak found
	// Note: The map is not needed to be returned, only the longest streak
	// as the problem statement only asks for the length of the longest consecutive sequence
	// and not the sequence itself.
	// The map is used to keep track of visited numbers to avoid counting them multiple times.
	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutiveSequence(nums)
	println(result) // Output: 4 (the longest consecutive sequence is [1, 2, 3, 4])

	nums2 := []int{0, 0, 1}
	result2 := longestConsecutiveSequence(nums2)
	println(result2) // Output: 2 (the longest consecutive sequence is [0, 1])

	nums3 := []int{1, 2, 0, 1}
	result3 := longestConsecutiveSequence(nums3)
	println(result3) // Output: 3 (the longest consecutive sequence is [0, 1, 2])

	nums4 := []int{1, 3, 5, 7}
	result4 := longestConsecutiveSequence(nums4)
	println(result4) // Output: 1 (the longest consecutive sequence is [1])

	nums5 := []int{1, 2, 3, 4, 5}
	result5 := longestConsecutiveSequence(nums5)
	println(result5) // Output: 5 (the longest consecutive sequence is [1, 2, 3, 4, 5])
}
