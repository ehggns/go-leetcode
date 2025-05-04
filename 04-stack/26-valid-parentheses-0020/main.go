package main

// validParentheses checks if a string of brackets is balanced.
// It returns true if every opening bracket has a corresponding closing bracket
// of the same type and all brackets are properly nested.
//
// The function takes a string s as input, which contains only the characters '(', ')', '{', '}', '[' and ']'.
//
// The function handles three types of brackets:
// - Parentheses: ( )
// - Square brackets: [ ]
// - Curly braces: { }
//
// It uses a stack to keep track of the opening brackets.
// When a closing bracket is encountered, it checks if it matches the last opening bracket.
// If it does, the last opening bracket is removed from the stack.
// If it doesn't match or if there are unmatched opening brackets left in the stack,
// the function returns false.
//
// Examples:
//   - validParentheses("()") returns true
//   - validParentheses("()[]{}") returns true
//   - validParentheses("(]") returns false
//   - validParentheses("([)]") returns false
//   - validParentheses("{[]}") returns true
//
// Time Complexity: O(n) where n is the length of the input string.
// Space Complexity: O(n) in the worst case when all characters are opening brackets.
func validParentheses(s string) bool {
	// Check if the length of the string is odd. If it is, return false immediately.
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s)) // Initializing an empty slice of runes to use as a stack

	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if _, exists := bracketMap[char]; exists {
			// If the character is a closing bracket, check if it matches the last opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false // Mismatch found
			}
			stack = stack[:len(stack)-1] // Pop the last opening bracket from the stack
		} else {
			stack = append(stack, char) // Append the opening bracket to the stack
		}
	}

	return len(stack) == 0 // If the stack is empty at the end, all brackets were matched correctly
}

func main() {
	// Example usage
	s := "{[()]}"
	result := validParentheses(s)
	println(result) // Output: true
	s2 := "{[(])}"
	result2 := validParentheses(s2)
	println(result2) // Output: false
	s3 := "{[}"
	result3 := validParentheses(s3)
	println(result3) // Output: false
	s4 := "{[()]}{}"
	result4 := validParentheses(s4)
	println(result4) // Output: true
	s5 := "{[()]}{"
	result5 := validParentheses(s5)
	println(result5) // Output: false
}
