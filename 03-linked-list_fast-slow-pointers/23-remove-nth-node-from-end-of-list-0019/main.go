package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of a linked list and returns the head of the modified list.
//
// The function uses a two-pointer technique (fast and slow) to find the node before the one to be removed:
// 1. A dummy node is created to handle edge cases when the head or tail needs to be removed.
// 2. The first pointer advances n+1 steps.
// 3. Both pointers then move forward simultaneously until the first pointer reaches the end.
// 4. At this point, the second pointer is at the node before the one to be removed.
// 5. The node is removed by updating the next pointer.
//
// Parameters:
//   - head: The head of the linked list.
//   - n: The position from the end of the list to remove (1-indexed).
//
// Returns:
//   - The head of the modified linked list.
//
// Time Complexity: O(L) where L is the length of the list.
// Space Complexity: O(1) as only constant extra space is used.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

func main() {
	// Example usage
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	newHead := removeNthFromEnd(head, n)

	// Print the modified list
	for newHead != nil {
		println(newHead.Val, " ")
		newHead = newHead.Next
	}
	// Output: 1 2 3 5
}
