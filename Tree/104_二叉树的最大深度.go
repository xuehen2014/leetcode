package Tree

/*
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 *
 */

//广度优先搜索
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxHeight int
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		//遍历该层
		for _, val := range curLevel {
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}
		maxHeight++
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return maxHeight
}

//深度优先搜索
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
