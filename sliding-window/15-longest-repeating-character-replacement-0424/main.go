package main

func LongestRepeatingCharacterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	count := make(map[byte]int)
	// maxFrequency is the count of the most frequent character in the current window
	start, maxLength, maxFrequency := 0, 0, 0

	for end := range len(s) {
		count[s[end]]++ // Increment the count of the current character
		maxFrequency = max(maxFrequency, count[s[end]])

		windowLength := end - start + 1
		if windowLength-maxFrequency > k { // if the window is invalid
			count[s[start]]-- // Decrease the count of the character at the start
			if count[s[start]] == 0 {
				delete(count, s[start]) // Remove it from the map if its count is zero
			}
			start++                        // Shrink the window from the left
			windowLength = end - start + 1 // Update the window length with the new start
		}
		maxLength = max(maxLength, windowLength) // Update the maximum length
	}
	return maxLength
}

func main() {
	// Example usage
	s := "AABABABBBA"
	k := 1 // Number of characters we can replace
	result := LongestRepeatingCharacterReplacement(s, k)
	println(result) // Output: 5
}
