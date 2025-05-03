package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList reverses a singly linked list and returns the new head.
//
// It iterates through the list once, maintaining three pointers:
// - prev: the node before the current node (initially nil)
// - curr: the current node being processed (initially the head)
// - next: the node after the current node (initially curr.Next)
//
// In each iteration, it reverses the Next pointer of curr to point to prev,
// advances prev to curr, and curr to next, thus reversing the list in-place.
//
// Time Complexity: O(n) where n is the number of nodes in the list
// Space Complexity: O(1) as only a constant amount of extra space is used
//
// Parameters:
//   - head: Pointer to the head of the linked list to reverse
//
// Returns:
//   - Pointer to the head of the reversed linked list (which was the tail of the original list)
func reverseList(head *ListNode) *ListNode {
	prev, curr, next := (*ListNode)(nil), head, head.Next

	println("Initial state: head:", head, "prev:", prev, "curr:", curr, "next:", next)

	for curr != nil {
		curr.Next = prev // Reverse the link
		prev = curr      // prev now points to the current node
		curr = next      // curr points to the next node
		if next != nil {
			next = next.Next // Move next to the next node if it exists
		}
		// Uncomment the following line to see the state of the list at each step
		println("Current:", curr, "Previous:", prev, "Next:", next)
	}
	// prev is now the new head of the reversed list and curr is nil
	head = prev
	return head
}

func main() {
	// Example usage
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reversedHead := reverseList(head)
	// Print the reversed list
	for node := reversedHead; node != nil; node = node.Next {
		println(node.Val)
	}
	// Output: 3 2 1
}
