package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If slow and fast pointers meet at some moment, there is a cycle in the linked list.
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	head := ListNode{Val: 0}
	curr := &head

	// Creating a linked list with 5 nodes
	for i := 1; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}

	hasCycle(&head) // Output: false
}
