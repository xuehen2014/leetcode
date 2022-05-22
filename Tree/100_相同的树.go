package Tree

/*
 * https://leetcode-cn.com/problems/same-tree/
 */

//深度优先搜索
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	leftRes := isSameTree(p.Left, q.Left)
	if leftRes == false {
		return false
	}
	rightRes := isSameTree(p.Right, q.Right)
	if rightRes == false {
		return false
	}
	return true
}

//广度优先搜索
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p==nil && q==nil {
		return true
	}
	if p==nil || q==nil {
		return false
	}
	var pCurLevel []*TreeNode
	var pNextLevel []*TreeNode
	var qCurLevel []*TreeNode
	var qNextLevel []*TreeNode

	pCurLevel = append(pCurLevel, p)
	qCurLevel = append(qCurLevel, q)
	for len(pCurLevel) > 0 && len(qCurLevel) > 0 {
		if len(pCurLevel) != len(qCurLevel) {
			return false
		}
		length := len(pCurLevel)
		for i:=0; i<length; i++ {
			node1 := pCurLevel[0]
			pCurLevel = pCurLevel[1:]
			node2 := qCurLevel[0]
			qCurLevel = qCurLevel[1:]

			if node1.Val != node2.Val {
				return false
			}
			if (node1.Left !=nil && node2.Left==nil) || (node1.Left== nil && node2.Left != nil ){
				return false
			}
			if (node1.Right !=nil && node2.Right==nil)|| (node1.Right== nil && node2.Right!= nil) {
				return false
			}
			if node1.Left !=nil {
				pNextLevel = append(pNextLevel, node1.Left)
			}
			if node1.Right != nil {
				pNextLevel = append(pNextLevel, node1.Right)
			}
			if node2.Left != nil {
				qNextLevel = append(qNextLevel, node2.Left)
			}
			if node2.Right != nil {
				qNextLevel = append(qNextLevel, node2.Right)
			}
		}
		pCurLevel =pNextLevel
		qCurLevel = qNextLevel
		pNextLevel = []*TreeNode{}
		qNextLevel = []*TreeNode{}
	}
	return len(pCurLevel) == 0 && len(qNextLevel) ==0
}
