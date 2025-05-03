package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle detects if a linked list has a cycle in it.
// It uses the Floyd's Cycle-Finding Algorithm (also known as "tortoise and hare" algorithm):
// - Two pointers, slow and fast, start at the head of the list.
// - The slow pointer moves one step at a time, while the fast pointer moves two steps.
// - If there's a cycle, the fast pointer will eventually meet the slow pointer.
// - If there's no cycle, the fast pointer will reach the end of the list.
//
// Parameters:
//   - head: The head node of the linked list.
//
// Returns:
//   - true if there is a cycle in the linked list.
//   - false if there is no cycle.
//
// Time Complexity: O(n), where n is the number of nodes in the linked list.
// Space Complexity: O(1), only constant extra space is used.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If slow and fast pointers meet at some moment, there is a cycle in the linked list
		// i.e. both variables are pointing to the same node address in memory.
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
	println("After creating a linked list with 5 nodes, the cycle check returns:")
	println(hasCycle(&head)) // Output: false

	// Creating a cycle for testing
	curr.Next = head.Next // Creating a cycle by pointing the last node to the second node
	hasCycle(&head) // Output: true
	println("After creating a cycle, the cycle check returns:")
	println(hasCycle(&head)) // Output: true
}
