package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoSortedLists merges two sorted linked lists into a single sorted list.
// It takes the heads of two sorted linked lists (l1 and l2) and returns the head
// of the merged sorted linked list.
//
// The function uses a dummy node approach to simplify handling the head of the result list.
// It iterates through both lists simultaneously, always selecting the smaller value to add
// to the result list. Once one list is exhausted, it appends the remainder of the other list.
//
// Time Complexity: O(n+m) where n and m are the lengths of l1 and l2 respectively.
// Space Complexity: O(1) as only constant extra space is used.
//
// Parameters:
//   - l1: head of the first sorted linked list
//   - l2: head of the second sorted linked list
//
// Returns:
//   - Head of the merged sorted linked list
func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy // This dummy node will help us easily return the head of the merged list

	// Traverse both lists and append the smaller value to the merged list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1 // Link the current node to l1
			l1 = l1.Next      // Move foward in l1
		} else {
			current.Next = l2 // Link the current node to l2
			l2 = l2.Next      // Move foward in l2
		}
		current = current.Next // Move to the next node in the merged list
	}

	// If one of the lists is not empty, append it to the merged list
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func main() {
	// Example usage
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList := mergeTwoSortedLists(l1, l2)

	// Print the merged list
	for mergedList != nil {
		println(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
	// Output: 1 1 2 3 4 4
}
