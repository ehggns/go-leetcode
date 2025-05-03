package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycleStart detects the starting node of a cycle in a linked list using
// Floyd's Tortoise and Hare algorithm.
// It returns the node where the cycle begins, or nil if no cycle exists.
//
// The algorithm uses two pointers moving at different speeds:
// 1. First, it detects if a cycle exists using slow and fast pointers
// 2. If a cycle is found, it resets one pointer to the head and moves both at the same speed
// 3. The point where they meet is the start of the cycle
//
//	The fast/slow pointer finds the meeting point in O(n) time, and then the second phase
//	also takes O(n) time to find the start of the cycle. Thus, the overall time complexity is O(n).
//
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1) as it only uses two pointers regardless of input size
func detectCycleStart(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head        // Reset slow pointer to head
	for slow != fast { // Move both pointers at the same speed. They will meet at the cycle start
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main() {
	// Example usage
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next // Creates a cycle

	cycleStart := detectCycleStart(head)
	if cycleStart != nil {
		println("Cycle starts at node with value:", cycleStart.Val) // 2
	} else {
		println("No cycle detected")
	}
}
