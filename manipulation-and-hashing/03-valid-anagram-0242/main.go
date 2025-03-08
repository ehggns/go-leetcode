/*
Brute force: Sort both strings and compare them.
Time complexity: O(nlogn)
Space complexity: O(1)

Optimized: Count the occurrences of each character in both strings and compare the counts.
Time complexity: O(n)
Space complexity: O(1)
*/
package main

import (
	"strings"
)

/*
Optimzed algorithm using a bucket slice to count the occurrences of each character, fit for ASCII
*/
func isAnagramBucket(s, t string) bool {
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	s = strings.ReplaceAll(s, " ", "")
	t = strings.ReplaceAll(t, " ", "")

	// Check if the lengths are equal before proceeding (early exit)
	if len(s) != len(t) {
		return false
	}

	// Create a slice to count the occurrences of each character in s (26 for lowercase letters)
	counts := make([]int, 26) // It is called a bucket because it is a fixed-size array that acts as a container for counting occurrences.

	// Count the occurrences of each character in s
	for _, char := range s {
		counts[char-'a']++ // ASCII. This creates a zero-based index where 'a'→0, 'b'→1, 'c'→2, etc. Incrementing the count at that index position
	}

	// Decrease the count for each character in t
	for _, char := range t {
		counts[char-'a']-- // Decrement the count for each character.
	}

	// Check if all counts are zero
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

/*
Optimzed algorithm that verifies if it is an anagram using a map to count characters and rune to handle unicode
*/
func isAnagramMap(s, t string) bool {
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	s = strings.ReplaceAll(s, " ", "")
	t = strings.ReplaceAll(t, " ", "")

	// Check if the lengths are equal before proceeding
	if len(s) != len(t) {
		return false
	}

	// Create a map to count the occurrences of each character in s
	counts := make(map[rune]int)

	// Count the occurrences of each character in s
	for _, char := range s {
		counts[char]++ // Increment the count for each character
	}

	// Decrease the count for each character in t
	for _, char := range t {
		counts[char]-- // Decrement the count for each character
	}

	// Check if all counts are zero
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	b := isAnagramMap(s, t)
	println(b) // Output: true

	s = "rat"
	t = "car"
	b = isAnagramMap(s, t)
	println(b) // Output: false

	s = "the cat"
	t = "the act"
	b = isAnagramMap(s, t)
	println(b) // Output: true

	s = "hello"
	t = "world"
	b = isAnagramMap(s, t)
	println(b) // Output: false

	s = "aabbcc"
	t = "abcabc"
	b = isAnagramBucket(s, t)
	println(b) // Output: true
}
