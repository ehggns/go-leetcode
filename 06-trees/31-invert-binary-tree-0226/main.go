package main

import "fmt"

// Binary Tree Node Definition
type BinTreeNode struct {
	Val   int
	Left  *BinTreeNode
	Right *BinTreeNode
}

// invertBinaryTree inverts a binary tree by swapping left and right children
// for each node in the tree. The inversion is performed recursively.
//
// Given a binary tree:
//
//	    1
//	   / \
//	  2   3
//	 / \ / \
//	4  5 6  7
//
// The inverted tree will be:
//
//	    1
//	   / \
//	  3   2
//	 / \ / \
//	7  6 5  4
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (for the recursion stack)
//
// Parameters:
//   - root: The root node of the binary tree to invert
//
// Returns:
//   - The root node of the inverted binary tree (same as input root)
//   - nil if the input root is nil
func invertBinaryTree(root *BinTreeNode) *BinTreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the left and right subtrees
	invertBinaryTree(root.Left)
	invertBinaryTree(root.Right)

	return root
}

func printTree(root *BinTreeNode) {
	printTreeHelper(root, "                        ", true)
}

func printTreeHelper(node *BinTreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}
	// Print right branch
	printTreeHelper(
		node.Right,
		prefix+(map[bool]string{true: "│   ", false: "    "})[isLeft],
		false)

	// Print current node
	branch := "└── "
	if !isLeft {
		branch = "┌── "
	}
	if prefix == "" {
		branch = ""
	}

	// Add color to the node value (green)
	nodeValue := fmt.Sprintf("\033[32m%d\033[0m", node.Val)

	// Print branches in blue
	coloredBranch := "\033[34m" + branch + "\033[0m"
	coloredPrefix := "\033[34m" + prefix + "\033[0m"

	if prefix == "" {
		fmt.Println(nodeValue)
	} else {
		fmt.Println(coloredPrefix + coloredBranch + nodeValue)
	}

	// Print left branch
	printTreeHelper(
		node.Left,
		prefix+(map[bool]string{true: "    ", false: "│   "})[isLeft],
		true)
}

func main() {
	// Example usage
	root := &BinTreeNode{Val: 7}
	root.Left = &BinTreeNode{Val: 6}
	root.Right = &BinTreeNode{Val: 3}
	root.Left.Left = &BinTreeNode{Val: 5}
	root.Left.Right = &BinTreeNode{Val: 4}
	root.Right.Left = &BinTreeNode{Val: 2}
	root.Right.Right = &BinTreeNode{Val: 1}
	root.Left.Left.Left = &BinTreeNode{Val: 8}
	root.Left.Left.Right = &BinTreeNode{Val: 9}
	root.Right.Left.Left = &BinTreeNode{Val: 10}
	root.Right.Left.Right = &BinTreeNode{Val: 11}
	root.Right.Right.Left = &BinTreeNode{Val: 12}
	root.Right.Right.Right = &BinTreeNode{Val: 13}
	root.Left.Right.Left = &BinTreeNode{Val: 14}
	root.Left.Right.Right = &BinTreeNode{Val: 15}

	fmt.Println("Original Binary Tree: root")
	printTree(root)

	invertedRoot := invertBinaryTree(root)

	fmt.Println("Inverted Binary Tree: root")
	printTree(invertedRoot)
}
