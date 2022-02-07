package Tree

/*
 *   https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
 */

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func nlevelOrder(root *NTreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var curLevel []*NTreeNode
	var nextLevel []*NTreeNode
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		var tmp []int
		for _, v := range curLevel {
			tmp = append(tmp, v.Val)
			for _, child := range v.Children {
				nextLevel = append(nextLevel, child)
			}
		}
		res = append(res, tmp)
		curLevel = nextLevel
		nextLevel = []*NTreeNode{}
	}
	return res
}
