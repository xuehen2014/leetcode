package Tree

/*
 * https://leetcode-cn.com/problems/maximum-width-of-binary-tree/
 * 二叉树的层序遍历
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodeWithID struct {
	Id int
	NodeInfo *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var curLevel  []TreeNodeWithID
	var nextLevel []TreeNodeWithID
	curLevel = append(curLevel, TreeNodeWithID{Id: 1, NodeInfo: root})
	for len(curLevel) > 0 {
		tmp := 0
		start := 0
		end := -1
		length := len(curLevel)
		for k, v := range curLevel {
			if k == 0 {
				start = v.Id
			}
			if k == length -1 {
				end = v.Id
			}
			if v.NodeInfo.Left != nil {
				nextLevel = append(nextLevel, TreeNodeWithID{Id: v.Id*2, NodeInfo: v.NodeInfo.Left})
			}
			if v.NodeInfo.Right != nil {
				nextLevel = append(nextLevel, TreeNodeWithID{Id: v.Id*2+1, NodeInfo: v.NodeInfo.Right})
			}
		}
		tmp = end-start+1
		if tmp > ans {
			ans = tmp
		}
		curLevel = nextLevel
		nextLevel = []TreeNodeWithID{}
	}
	return ans
}
