package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper function to merge two sorted linked lists to be used in the mergeKListsRecursive function
func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummy.Next
}

// mergeKListsRecursive merges k sorted linked lists into a single sorted linked list.
// It uses a divide-and-conquer approach to recursively merge the lists:
//   - If the input array is empty, it returns nil.
//   - If there's only one list, it returns that list.
//   - If there are two lists, it directly merges them using mergeTwoSortedLists.
//   - For more than two lists, it splits the array in half, recursively merges each half,
//     and then combines the results.
//
// Time Complexity: O(N log k) where N is the total number of nodes across all lists and k is the number of lists.
// Space Complexity: O(log k) for the recursion stack.
//
// Explaning the Time Complexity:
// - Each merge operation takes O(N) time, where N is the total number of nodes in the two lists being merged.
// - The number of merge operations is log(k) because we are halving the number of lists at each step.
// - Therefore, the overall time complexity is O(N log k).
//
// Parameters:
//   - lists: A slice of pointers to the head nodes of k sorted linked lists.
//
// Returns:
//   - A pointer to the head of the merged sorted linked list.
func mergeKListsRecursive(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoSortedLists(lists[0], lists[1])
	}

	mid := len(lists) / 2
	left := mergeKListsRecursive(lists[:mid])
	right := mergeKListsRecursive(lists[mid:])
	return mergeTwoSortedLists(left, right)
}

func main() {
	// Example usage
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := &ListNode{2, &ListNode{6, nil}}
	lists := []*ListNode{l1, l2, l3}

	mergedList := mergeKListsRecursive(lists)

	// Print the merged list
	for mergedList != nil {
		println(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
	// Output: 1 1 2 3 4 4 5 6
}
