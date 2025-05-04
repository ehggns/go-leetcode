# Go LeetCode Coding Techniques

This document outlines the key coding techniques, algorithms, data structures, and Go-specific idioms used throughout the LeetCode problem solutions in this repository.

## Table of Contents

1. [Data Structures](#data-structures)
2. [Two-Pointer Techniques](#two-pointer-techniques)
3. [Sliding Window Patterns](#sliding-window-patterns)
4. [Hash-Based Techniques](#hash-based-techniques)
5. [Linked List Manipulation](#linked-list-manipulation)
6. [Fast & Slow Pointer Technique](#fast--slow-pointer-technique)
7. [Divide-and-Conquer](#divide-and-conquer)
8. [Go-Specific Idioms](#go-specific-idioms)
9. [Time and Space Complexity Analysis](#time-and-space-complexity-analysis)

## Data Structures

### Maps (Hash Maps)

Used extensively for O(1) lookups:

```go
// Creating a map to store elements for quick lookups
seen := make(map[int]bool)
// or
counts := make(map[int]int) // key is element, value is frequency
```

Example usage in problems:

- Two Sum (Problem 1)
- Contains Duplicate (Problem 217)
- Valid Anagram (Problem 242)

### Slices

Used for dynamic lists and arrays:

```go
// Creating a slice with initial capacity
result := make([]int, 0, capacity)

// Appending elements
result = append(result, value)
```

Example usage in problems:

- Top K Frequent Elements (Problem 347)
- Product of Array Except Self (Problem 238)

### Map as Set Implementation

Since Go has no built-in set data structure, maps are commonly used with empty structs as values:

```go
// Creating a set using map with empty struct (memory efficient)
set := make(map[int]struct{})

// Adding an element to the set
set[item] = struct{}{}

// Checking if an element exists in the set
if _, exists := set[item]; exists {
    // Element found
}
```

Example usage in problems:

- Contains Duplicate (Problem 217)
- Group Anagrams (Problem 49)

## Two-Pointer Techniques

### Basic Two-Pointer Approach

Used when working with sorted arrays or when processing elements from both ends:

```go
left, right := 0, len(array)-1
for left < right {
    // Process elements
    if condition {
        left++
    } else {
        right--
    }
}
```

Example usage in problems:

- Valid Palindrome (Problem 125)
- Two Sum II - Input Array is Sorted (Problem 167)
- Container With Most Water (Problem 11)

### Fast & Slow Pointers

Used particularly for cycle detection and finding middle elements:

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next          // Move one step
    fast = fast.Next.Next     // Move two steps

    if slow == fast {
        // Cycle detected!
    }
}
```

Example usage in problems:

- Middle of the Linked List (Problem 876)
- Linked List Cycle (Problem 141)
- Linked List Cycle II (Problem 142)

## Sliding Window Patterns

### Fixed Window Size

Used when processing subarrays of fixed length:

```go
// Calculate initial window
sum := 0
for i := 0; i < k; i++ {
    sum += nums[i]
}

maxSum := sum
// Slide the window
for i := k; i < len(nums); i++ {
    sum = sum - nums[i-k] + nums[i]  // Remove leftmost, add rightmost
    maxSum = max(maxSum, sum)
}
```

Example usage in problems:

- Maximum Average Subarray I (Problem 643)

### Dynamic Window Size

Used when the window size can vary based on certain conditions:

```go
start, end := 0, 0
for end < len(s) {
    // Expand window by including the character at end
    // ...

    for {condition} {
        // Contract window from the left
        // ...
        start++
    }
    end++
}
```

Example usage in problems:

- Longest Substring Without Repeating Characters (Problem 3)
- Minimum Window Substring (Problem 76)
- Longest Repeating Character Replacement (Problem 424)

## Hash-Based Techniques

### Frequency Counting

Used to count occurrences of elements:

```go
counts := make(map[int]int)
for _, num := range nums {
    counts[num]++
}
```

Example usage in problems:

- Valid Anagram (Problem 242)
- Top K Frequent Elements (Problem 347)

### Character Frequency for Strings

Using arrays for character frequency counts (more efficient for alphabet):

```go
// For lowercase English letters
counts := make([]int, 26)
for _, char := range s {
    counts[char-'a']++
}
```

Example usage in problems:

- Valid Anagram (Problem 242)
- Group Anagrams (Problem 49)

### Map for Grouping

Using maps to group related elements:

```go
groups := make(map[string][]string)
for _, item := range items {
    key := getKey(item)
    groups[key] = append(groups[key], item)
}
```

Example usage in problems:

- Group Anagrams (Problem 49)

## Linked List Manipulation

### Dummy Node Technique

Using a dummy node to simplify linked list operations:

```go
dummy := &ListNode{}
current := dummy
// Perform operations
// ...
return dummy.Next  // The actual head of the result list
```

Example usage in problems:

- Merge Two Sorted Lists (Problem 21)
- Merge k Sorted Lists (Problem 23)

### Reversing Linked Lists

In-place reversal using three pointers:

```go
prev, curr, next := (*ListNode)(nil), head, head.Next
for curr != nil {
    curr.Next = prev  // Reverse the link
    prev = curr       // Move prev forward
    curr = next       // Move curr forward
    if next != nil {
        next = next.Next  // Move next forward
    }
}
return prev  // New head
```

Example usage in problems:

- Reverse Linked List (Problem 206)
- Reorder List (Problem 143)

## Fast & Slow Pointer Technique

### Finding Middle of Linked List

Using fast and slow pointers to find the middle node:

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
// slow is at the middle
```

Example usage in problems:

- Middle of the Linked List (Problem 876)
- Reorder List (Problem 143)

### Cycle Detection

Floyd's Tortoise and Hare algorithm:

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
        return true  // Cycle detected
    }
}
return false  // No cycle
```

Example usage in problems:

- Linked List Cycle (Problem 141)
- Linked List Cycle II (Problem 142)

## Divide-and-Conquer

Divide and conquer is an algorithmic paradigm where a problem is broken down into smaller subproblems, which are solved independently, and their solutions are combined to solve the original problem.

### Recursive Divide-and-Conquer

This pattern involves:

1. Dividing the problem into smaller subproblems
2. Solving the subproblems recursively
3. Combining the results to form a solution to the original problem

```go
func divideAndConquer(problem []Type) Result {
    // Base cases
    if len(problem) <= threshold {
        return solveDirectly(problem)
    }

    // Divide
    mid := len(problem) / 2
    leftResult := divideAndConquer(problem[:mid])  // First half
    rightResult := divideAndConquer(problem[mid:]) // Second half

    // Combine
    return combine(leftResult, rightResult)
}
```

Example usage in problems:

- Merge K Sorted Lists (Problem 23)
- Sort an Array (Problem 912)
- Maximum Subarray (Problem 53)

### Slice Operations in Go for Divide-and-Conquer

Go's slice expressions make implementing divide-and-conquer approaches elegant:

```go
// Create a slice with elements from index 0 to mid-1
leftHalf := array[:mid]

// Create a slice with elements from index mid to end
rightHalf := array[mid:]
```

Key properties of slice operations:

- O(1) time complexity (no copying of elements)
- They create a view into the underlying array
- Changes to elements in the slice affect the original array

### Example: Merge K Sorted Lists

The "Merge K Sorted Lists" problem demonstrates a classic divide-and-conquer approach:

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    // Divide
    mid := len(lists) / 2
    left := mergeKLists(lists[:mid])   // Process first half
    right := mergeKLists(lists[mid:])  // Process second half

    // Combine (using a helper function to merge two lists)
    return mergeTwoLists(left, right)
}
```

This achieves O(N log k) time complexity, where N is the total number of nodes across all lists and k is the number of lists. The logarithmic factor comes from the divide-and-conquer approach which reduces the problem size by half in each recursive call.

## Go-Specific Idioms

### Empty Struct as Map Value

Using empty structs to save memory in sets:

```go
set := make(map[string]struct{})
set["key"] = struct{}{}  // Zero memory allocation for the value
```

### Map Existence Check

Idiomatic way to check if a key exists in a map:

```go
if val, exists := myMap[key]; exists {
    // Key exists, and val is its value
}
```

### Type Conversion for Runes

Converting between rune/byte and their corresponding indices:

```go
index := char - 'a'  // Convert character to zero-based index (0-25)
```

### Nil as Zero Value for Pointers

Using nil as the zero value for pointer types:

```go
var prev *ListNode = nil  // Explicit nil initialization
// or
prev := (*ListNode)(nil)  // Type conversion to explicitly indicate nil pointer
```

## Time and Space Complexity Analysis

### Time Complexity Notation

- O(1) - Constant time
- O(log n) - Logarithmic time
- O(n) - Linear time
- O(n log n) - Linearithmic time
- O(n²) - Quadratic time

### Space Complexity Considerations

- O(1) - Constant space (no additional data structures that scale with input)
- O(n) - Linear space (data structures that grow proportionally with input size)

### Complexity Analysis Examples

- Two pointer techniques typically use O(n) time and O(1) space
- Hash-based solutions typically use O(n) time and O(n) space
- Sorting-based solutions typically use O(n log n) time

---

This document provides an overview of the main coding techniques used throughout this repository. Each problem solution contains specific comments detailing the approach, time and space complexity, and any special considerations for that particular problem.
