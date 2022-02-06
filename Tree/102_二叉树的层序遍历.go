package Tree

/*
 *  二叉树的层序遍历
 *
 *  https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var res [][]int
	queue = append(queue, root)
	for len(queue) > 0 {
		currentLength := len(queue)
		var tmp []int
		for ; currentLength>0; currentLength-- {
			v := queue[0]
			queue = queue[1:]
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		var tmp []int
		for _, val := range curLevel {
			tmp = append(tmp, val.Val)
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}
		res = append(res, tmp)
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return res
}


func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var res [][]int
	var tmp []int
	curNum, nextLevelNum := 1, 0
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp,  node.Val)
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
	return res
}



