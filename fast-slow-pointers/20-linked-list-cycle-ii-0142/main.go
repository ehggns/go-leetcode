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
//
// 1. First, it detects if a cycle exists using slow and fast pointers
//
// 2. If a cycle is found, it resets one pointer to the head and moves both at the same speed
//
// 3. The point where they meet is the start of the cycle
//
// The fast/slow pointer finds the meeting point in O(n) time, and then the second phase
// also takes O(n) time to find the start of the cycle. Thus, the overall time complexity is O(n).
//
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1) as it only uses two pointers regardless of input size
//
// Mathematical proof:
//
// 1. Let's denote:
//    - k = distance from head to the start of the cycle
//    - m = distance from the cycle start to the meeting point
//    - c = length of the cycle
//
// 2. When slow and fast pointers meet:
//    - Slow pointer traveled distance: k + m
//    - Fast pointer traveled distance: k + m + n*c (where n is some integer)
//
// 3. Since the fast pointer moves twice as fast as the slow pointer:
//    - 2*(k + m) = k + m + n*c, then
//    - k + m = n*c, then
//    - k = n*c - m
//
// 4. This means k is congruent to -m modulo c.
//		In other words, k is equivalent to -m in the context of modulo c.
//
// 		In mathematical notation, this would be written as: k ≡ -m (mod c).
//    This means that when k and -m are divided by c, they yield the same remainder.
//
// 		The definition for modulo to integer values is: `a % b = a - (⌊a/b⌋ × b)`
//
// 		e.g. for -8 % 12, we can calculate it as:
//
//    -8 - (⌊-8/12⌋ × 12) = -8 - (-1 × 12) = -8 + 12 = 4
//
//    So -8 % 12 = 4, which means that -8 and 4 are congruent modulo 12.
//
// 		In other words, -8 and 4 are in the same equivalence class under modulo 12:
//    -8 ≡ 4 (mod 12).
//
//    A quick example:
//    from the algebraic definition of division: a = b * q + r
//    where a is the dividend, b is the divisor, q is the quotient, and r is the remainder.
//    For example, if we take 13 and divide it by 4:
//    13 = 3*4 + 1   →   13 ≡ 1 (mod 4)
//    15 = 4*4 - 1   →   15 ≡ -1 (mod 4)
//    15 = 3*4 + 3   →   15 ≡ 3 (mod 4)
//    Therefore, 13 and 15 are congruent modulo 4.
//
//    The distance from head to cycle start (k) is the same as the distance from
//    the meeting point to cycle start if we travel in the cycle.
//
//    So k is both the distance from the head to the start of the cycle and the distance from
//    the meeting point to the start of the cycle.
//
//    Since they are congruent modulo c, we can say that k is equivalent to -m in the context of the cycle,
//		which means that the distance from the meeting point to the start of the cycle is also k.
//
// 5. Therefore, if we reset one pointer to the head and move both pointers at the same speed,
//    they will meet at the start of the cycle after k steps.
//
// 6. If there is no cycle, the fast pointer will reach the end of the list, and we return nil.
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
