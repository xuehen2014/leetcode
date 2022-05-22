package Tree

/*
 *  https://leetcode-cn.com/problems/symmetric-tree/
 */

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p ==nil && q == nil {
		return true
	}
	if p==nil || q==nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	res1 := check(p.Left, q.Right)
	if !res1 {
		return false
	}
	res2 := check(p.Right, q.Left)
	if !res2 {
		return false
	}
	return true
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var curLevel  []*TreeNode
	var nextLevel []*TreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		for _, node := range curLevel {
			if node != nil {
				nextLevel = append(nextLevel, node.Left)
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if !isSymmetricSlice(curLevel) {
			return false
		}
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return true
}

func isSymmetricSlice(list []*TreeNode) bool {
	l :=0
	r := len(list) - 1
	for l<=r {
		if (list[l] !=nil &&list[r] == nil ) || (list[l]==nil && list[r] != nil) {
			return false
		}
		if list[l] !=nil && list[r] != nil && list[l].Val != list[r].Val {
			return false
		}
		l++
		r--
	}
	return true
}