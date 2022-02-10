package Tree

/*
 *  https://leetcode-cn.com/problems/invert-binary-tree/
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		for _, val := range curLevel {
			val.Left, val.Right = val.Right, val.Left  //翻转
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}

		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return root
}


func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}


