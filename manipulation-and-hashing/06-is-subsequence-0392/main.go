/*
isSubsequence determines whether string s is a subsequence of string t.
A subsequence is a sequence that can be derived from another sequence by
deleting some or no elements without changing the order of the remaining elements.

For example, "ace" is a subsequence of "abcde" while "aec" is not.

Parameters:
  - s: the potential subsequence string
  - t: the source string to check against

Returns:
  - true if s is a subsequence of t, or if s is empty
  - false if s is not a subsequence of t, or if t is empty but s is not

Time Complexity: O(n), where n is the length of t
Space Complexity: O(1), only constant extra space is used
*/
package main

// isSubsequence checks if s is a subsequence of t using two pointers.
func isSubsequence(s, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen == 0 {
		return true
	}
	if tLen == 0 {
		return false
	}

	sIndex, tIndex := 0, 0 // Initialize two pointers for s and t
	// Iterate through t and check if characters in s match
	for sIndex < sLen && tIndex < tLen {
		if s[sIndex] == t[tIndex] {
			sIndex++ // Only increment sIndex if characters match
		}
		tIndex++ // Always increment tIndex
	}

	return sIndex == sLen // If sIndex equals sLen, all characters in s were found in t
}

func main() {
	// Example usage
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	println(result) // Output: true
	s = "axc"
	t = "ahbgdc"
	result = isSubsequence(s, t)
	println(result) // Output: false
	s = ""
	t = "ahbgdc"
	result = isSubsequence(s, t)
	println(result) // Output: true
	s = "abc"
	t = ""
	result = isSubsequence(s, t)
	println(result) // Output: false
	s = "abc"
	t = "abc"
	result = isSubsequence(s, t)
	println(result) // Output: true
	s = "abc"
	t = "ab"
	result = isSubsequence(s, t)
	println(result) // Output: false
}