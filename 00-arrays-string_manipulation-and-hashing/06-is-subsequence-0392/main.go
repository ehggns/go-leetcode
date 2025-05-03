package main

/*
isSubsequence determines if string s is a subsequence of string t.
A subsequence is a sequence that can be derived from another sequence by deleting some
or no elements without changing the order of the remaining elements.

Parameters:
  - s: the string to check if it is a subsequence
  - t: the string to check against

Returns:
  - bool: true if s is a subsequence of t, false otherwise

Example:

	isSubsequence("abc", "ahbgdc") returns true
	isSubsequence("axc", "ahbgdc") returns false
*/
func isSubsequence(s string, t string) bool {
	sIndex := 0
	tIndex := 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++ // Always increment tIndex
	}

	// If we have matched all characters of s, return true
	return sIndex == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	println(result) // Output: true

	s = "axc"
	t = "ahbgdc"
	result = isSubsequence(s, t)
	println(result) // Output: false
}
