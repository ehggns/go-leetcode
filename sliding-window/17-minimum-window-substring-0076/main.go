package main

// minimumWindowSubstring finds the shortest substring in s that contains all characters from t.
//
// Given two strings s and t, this function returns the minimum window substring of s such that
// every character in t (including duplicates) is included in the window. If there is no such
// substring, it returns an empty string "".
//
// The algorithm uses a sliding window approach with two pointers:
// 1. Expand the window by moving the end pointer to include required characters
// 2. Contract the window from the left while maintaining all required characters
// 3. Track the minimum valid window seen so far
//
// Time Complexity: O(|s| + |t|) where |s| and |t| are the lengths of strings s and t
// Space Complexity: O(|s| + |t|) for the character frequency maps
//
// Parameters:
//   - s: The source string to search in
//   - t: The target string containing characters that must be included
//
// Returns:
//   - The minimum window substring that contains all characters in t, or an empty string if not found
func minimumWindowSubstring(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	countT := make(map[byte]int)
	for i := range len(t) {
		countT[t[i]]++
	}

	countS := make(map[byte]int) // Frequency map for characters in s
	start, end, minLen := 0, 0, len(s)+1 // Initialize minLen to a value larger than any possible window
	minStart := 0 // Start index of the minimum window
	required := len(countT) // Number of unique characters in t that need to be present in the window
	formed := 0 // Number of unique characters in the current window that match the required count

	// Expand the window by moving the end pointer to include required characters
	for end < len(s) {
		char := s[end]
		countS[char]++

		// Check if the current character is part of t and if we have enough of it
		if countT[char] > 0 && countS[char] == countT[char] {
			formed++
		}

		// Contract the window from the left while maintaining all required characters
		// If we have all required characters, try to minimize the window size
		for start <= end && formed == required {
			if end-start+1 < minLen { // end-start+1 is the current window size
				minLen = end - start + 1 // Update minLen if we found a smaller window
				minStart = start // Update minStart to the current start index
			}

			countS[s[start]]-- // Remove the character at the start of the window

			// If the character at start is part of t and we have less than required, reduce formed
			// This means we can no longer form a valid window
			if countT[s[start]] > 0 && countS[s[start]] < countT[s[start]] {
				formed-- // Reduce formed count since we no longer have enough of this character
			}
			start++ // Move the start pointer to the right to try and find a smaller window
		}
		end++ // Move the end pointer to the right to expand the window
	}

	if minLen == len(s)+1 {
		return ""
	}
	// s[minStart : minStart+minLen] is a substring of s starting at minStart with length minLen
	return s[minStart : minStart+minLen] // Return the minimum window substring
}

func main() {
	// Example usage
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minimumWindowSubstring(s, t)
	println(result) // Output: "BANC"
}