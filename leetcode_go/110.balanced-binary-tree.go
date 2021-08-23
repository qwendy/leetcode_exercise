/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (41.97%)
 * Likes:    1525
 * Dislikes: 139
 * Total Accepted:    368K
 * Total Submissions: 876.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 *
 *
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package practice

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := treeHeight(root.Left)
	right := treeHeight(root.Right)
	if left-right > 1 || right-left > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)

}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeHeight(root.Left)
	right := treeHeight(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// @lc code=end
