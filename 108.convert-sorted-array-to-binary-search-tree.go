/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (53.32%)
 * Likes:    1502
 * Dislikes: 154
 * Total Accepted:    310.7K
 * Total Submissions: 582.7K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNozde struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package practice

func sortedArrayToBST(nums []int) *TreeNode {
	head := &TreeNode{}
	for i := 0; i < len(nums); i++ {
		head = insertBinarySearchTree(nums[i], head)
	}
	return head
}
func insertBinarySearchTree(value int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value}
	}
	if value < root.Val {
		root = insertBinarySearchTree(value, root.Left)
	}
	if value > root.Val {
		root = insertBinarySearchTree(value, root.Right)
	}
	if root.Left != nil && root.Left.Left != nil {
		root = rightRotate(root)
	}
	if root.Right != nil && root.Right.Right != nil {
		root = leftRotate(root)
	}
	if root.Left != nil && root.Left.Right != nil {
		root.Left.Left = root.Left.Right
		root.Left.Right = nil
	}
	if root.Right != nil && root.Right.Left != nil {
		root.Right.Right = root.Right.Left
		root.Right.Left = nil
	}
	return root
}

func rightRotate(root *TreeNode) *TreeNode {
	x := root.Left
	root.Left = x.Right
	x.Right = root
	return x
}

func leftRotate(root *TreeNode) *TreeNode {
	x := root.Right
	root.Right = x.Left
	x.Left = root
	return x
}

// @lc code=end
