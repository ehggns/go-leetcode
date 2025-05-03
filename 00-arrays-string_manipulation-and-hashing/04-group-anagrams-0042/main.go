/*
This package implements a solution for grouping anagrams.

An anagram is a word formed by rearranging the letters of another word,
using all the original letters exactly once. For example, "eat", "tea", and "ate"
are anagrams of each other.

The solution uses a character frequency encoding approach to efficiently
identify anagrams. Each string is encoded into a representation where each
character in the alphabet is followed by its frequency in the string.
This encoded string serves as a unique key for anagram groups.

Example:

	Input: ["eat", "tea", "tan", "ate", "nat", "bat"]
	Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]

Time complexity: O(n*k) where n is the number of strings and k is the
maximum length of a string.
Space complexity: O(n*k) for storing the result.
*/
package main

/*
getEncodedString takes a string and returns an encoded representation where each letter
is followed by its frequency count in the input string.

The encoding works by:
1. Creating a frequency map of 26 lowercase letters
2. Converting the frequency map into a string where each letter is followed by its count

For example, "hello" becomes "h1e1l2o1..." (showing only relevant letters)

Parameters:
  - s: The input string to be encoded

Returns:

	A string where each letter a-z is followed by its frequency in the input

Time complexity: O(n) where n is the length of input string
Space complexity: O(1) as we use fixed size arrays
*/
func getEncodedString(s string) string {
	frequencyList := make([]int, 26)
	for _, char := range s {
		frequencyList[char-'a']++
	}

	encodedString := make([]byte, 0, 26)
	for i, count := range frequencyList {
		encodedString = append(encodedString, byte(i+'a'), byte(count+'0'))
	}
	return string(encodedString)
}

/*
groupAnagrams groups a slice of strings into anagram groups and returns them as a slice of string slices.
An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
using all the original letters exactly once.

Parameters:
  - strs: A slice of strings to be grouped into anagrams

Returns:
  - [][]string: A slice where each inner slice contains strings that are anagrams of each other.
    Returns an empty slice if the input is empty.

Example:

	Input: ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	if len(strs) == 0 {
		return [][]string{}
	}

	for _, str := range strs {
		encodedStr := getEncodedString(str)                  // Use the encoded string as a key
		groups[encodedStr] = append(groups[encodedStr], str) // Append the original string to the group
	}

	// Convert the map to a slice of slices
	result := make([][]string, 0, len(groups)) // Preallocate the slice
	for _, words := range groups {             // Iterate over the map values
		result = append(result, words) // Append each group to the result
	}

	return result
}

func main() {
	// Example usage
	encodedStr := getEncodedString("anagram")
	println("Encoded string for the word \"anagram\" is:", encodedStr) // Output: a3b0c0d0e0f0g1h0i0j0k0l0m1n1o0p0q0r1s0t0u0v0w0x0y0z0

	// Example usage of groupAnagrams
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupedAnagrams := groupAnagrams(words)

	// Print the result
	println("Result of set of words [\"eat\", \"tea\", \"tan\", \"ate\", \"nat\", \"bat\"]:")
	for i, group := range groupedAnagrams {
		print("Group ", i, ": ")
		for j, word := range group {
			if j > 0 {
				print(" ")
			}
			print(word)
		}
		println("")
	}
}
