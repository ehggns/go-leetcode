package main

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
//
// The function uses a sliding window approach with two pointers:
// - A 'start' pointer marking the beginning of the current substring
// - An 'end' pointer marking the end of the current substring
//
// It maintains a set of characters in the current window to quickly check for duplicates.
// When a duplicate is found, the window is shrunk from the left until the duplicate is removed.
//
// Parameters:
//   - s: The input string to be processed
//
// Returns:
//   - The length of the longest substring without repeating characters
//
// Time Complexity: O(n) where n is the length of the input string
// Space Complexity: O(min(m, n)) where m is the size of the character set
func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]struct{})
	start, end, maxLength := 0, 0, 0

	for end < len(s) {
		if _, exists := charSet[s[end]]; !exists {
			charSet[s[end]] = struct{}{}
			end++
			maxLength = max(maxLength, end-start)
		} else {
			delete(charSet, s[start])
			start++
		}
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	println(result) // Output: 3
}