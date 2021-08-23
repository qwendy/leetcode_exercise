/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (48.70%)
 * Likes:    885
 * Dislikes: 167
 * Total Accepted:    256.3K
 * Total Submissions: 526.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 *
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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var s [][]int
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		levelResult := make([]int, 0)
		var newStack []*TreeNode
		for i := 0; i < len(stack); i++ {
			levelResult = append(levelResult, stack[i].Val)
			if stack[i].Left != nil {
				newStack = append(newStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				newStack = append(newStack, stack[i].Right)
			}
		}
		s = append(s, levelResult)
		stack = newStack
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-1-i], s[i]
	}
	return s
}

// @lc code=end
