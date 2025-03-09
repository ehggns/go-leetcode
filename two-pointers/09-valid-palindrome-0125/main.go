package main

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

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
