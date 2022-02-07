package Tree

/*
 *  https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
 *  二叉树的层序遍历
 */

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var ans []float64
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		var tmp int
		curLevelSize := len(curLevel)
		for _, val := range curLevel {
			tmp = tmp+val.Val
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}
		avgVal := float64(tmp)/float64(curLevelSize)
		ans = append(ans, avgVal)
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return ans
}