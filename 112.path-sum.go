/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (39.03%)
 * Likes:    1202
 * Dislikes: 380
 * Total Accepted:    367.4K
 * Total Submissions: 941.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 *
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return calSum(0, sum, root)
}

func calSum(curSum, sum int, root *TreeNode) bool {
	if root == nil {
		return false
	}
	curSum = root.Val + curSum
	if curSum == sum && root.Left == nil && root.Right == nil {
		return true
	}
	if calSum(curSum, sum, root.Left) {
		return true
	}
	return calSum(curSum, sum, root.Right)
}

// @lc code=end
