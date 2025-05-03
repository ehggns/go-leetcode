package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList rearranges a singly linked list in-place to follow the pattern:
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...
//
// It modifies the list by reordering the nodes without using extra space for a new list.
// Interleaving the nodes from the first half and the reversed second half of the list.
//
// The algorithm works in three steps:
// 1. Find the middle of the linked list using fast and slow pointers
// 2. Reverse the second half of the linked list
// 3. Merge the first half and the reversed second half by interleaving nodes
//
// Time Complexity: O(n) where n is the number of nodes in the list
// Space Complexity: O(1) as it only uses a constant amount of extra space
//
// Parameters:
//   - head: Pointer to the head node of the linked list
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find the middle of the list using the fast-slow pointer technique
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Step 2: Reverse the second half of the list
	prev, curr, next := (*ListNode)(nil), slow, slow.Next
	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = next
		if next != nil {
			next = next.Next
		}
	}
	// Step 3: Merge the two halves
	first, second := head, prev // prev is now the head of the reversed second half
	for second.Next != nil {    // We stop when we reach the end of the second half
		// Store the next nodes
		tmp1, tmp2 := first.Next, second.Next
		// Rearrange pointers
		first.Next = second
		second.Next = tmp1
		// Move to the next nodes
		first, second = tmp1, tmp2
	}
}

func main() {
	// Example usage
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	// Print the reordered list
	for node := head; node != nil; node = node.Next {
		println(node.Val)
	}
	// Output: 1 4 2 3
}
