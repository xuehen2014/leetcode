package Tree

/*
 * https://leetcode-cn.com/problems/deepest-leaves-sum/
 * 二叉树的层序遍历
 */

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		var tmp int
		for _, val := range curLevel {
			tmp = tmp+val.Val
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}
		res = tmp
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return res
}

