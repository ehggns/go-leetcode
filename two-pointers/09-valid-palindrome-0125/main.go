package main

// isAlphaNumeric verifies if a character (a byte) is alphanumeric and returns a boolean
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

/*
isPalindrome verifies if an input string is a palindrome
(i.e. the same characters from the left to right and right to left)
It uses the pattern of two pointers, each one starting from an end of the string.
While the left pointer does not meet the right pointer, it checks if the characters
are the same from both pointers. If they are not equal, it returns false. While
they are equal, left pointer advances forward, and right pointer go backward.
When l == r, the loop stops and it returns true.
This approach achieves O(n) time complexity with O(1) space complexity.
*/
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNumeric(s[l]) {
			l++
		}

		for l < r && !isAlphaNumeric(s[r]) {
			r--
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func main() {
	str1 := "race car"
	b := isPalindrome(str1)
	println(b) // expects true

	str2 := "race a car"
	b = isPalindrome(str2)
	println(b) // expects false
}
