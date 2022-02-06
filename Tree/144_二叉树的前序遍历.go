package Tree

/*
 * 144-> https://leetcode-cn.com/problems/binary-tree-preorder-traversal
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归前序遍历
func preorderTraversal1(root *TreeNode) []int {
	var resultList []int
	if root == nil {
		return resultList
	}
	preTraversal(root, &resultList)
	return resultList
}

func preTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preTraversal(root.Left, res)
	preTraversal(root.Right, res)
}

//非递归前序遍历
func preorderTraversal(root *TreeNode) []int {
	var resultList []int
	if root == nil {
		return resultList
	}
	stackList := newMyStack()
	for root != nil || stackList.length() > 0 {
		for root != nil {
			resultList = append(resultList, root.Val)
			stackList.push(root)
			root = root.Left
		}
		node := stackList.pop()
		root = node.Right
	}

	return resultList
}

