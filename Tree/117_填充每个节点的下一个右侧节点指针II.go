package Tree

/*
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
 * 二叉树的层序遍历
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var curLevel []*Node
	var nextLevel []*Node
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		for _, val := range curLevel {
			if val.Left != nil {
				nextLevel = append(nextLevel, val.Left)
			}
			if val.Right != nil {
				nextLevel = append(nextLevel, val.Right)
			}
		}
		nextLevelLength := len(nextLevel)
		for i := 0; i < nextLevelLength-1; i++ {
			nextLevel[i].Next = nextLevel[i+1]
		}
		curLevel = nextLevel
		nextLevel = []*Node{}
	}
	return root
}
