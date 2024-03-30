package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

func main() {
	tree := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3, Right: &TreeNode{
			Val: 4,
		}},
	}

	fmt.Println(searchBST(tree, 3))
}
