package main

// ListNode represents a node in a singly-linked list.
// Each node contains a value (Val) and a pointer to the next node in the list (Next).
// A nil Next pointer indicates the end of the list.
// In coding exercises, a more idiomatic approach is to tailor the data structure to the problem.
// In this case, we are using a singly-linked list to represent the nodes.
type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode returns the middle node of a linked list.
// If the list has an even number of nodes, it returns the second middle node.
//
// It uses the fast/slow pointer technique:
// - slow pointer moves one step at a time
// - fast pointer moves two steps at a time
// - when fast reaches the end, slow is at the middle
//
// Time Complexity: O(n), where n is the number of nodes in the list
// Space Complexity: O(1)
//
// Parameters:
//   - head: The head node of the linked list
//
// Returns:
//   - The middle node, or nil if the list is empty
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	head := ListNode{Val: 0}
	curr := &head

	// Creating a linked list with 5 nodes
	for i := 1; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}

	middle := middleNode(&head)
	println(middle.Val) // Output: 3
}
