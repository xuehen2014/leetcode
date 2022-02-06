package Tree

/*
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
 */

//递归后序遍历
func postorderTraversal(root *TreeNode) []int {
	var res []int
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traversal(root.Left, res)
	}
	if root.Right != nil {
		traversal(root.Right, res)
	}
	*res = append(*res, root.Val)
	return
}

//非递归后续遍历
func postorderTraversal2(root *TreeNode) []int {
	var res []int
	myStack := newMyStack()
	var preNode *TreeNode
	for root != nil || myStack.length() > 0 {
		for root != nil {
			myStack.push(root)
			root = root.Left
		}
		root = myStack.peek()
		if root.Right == nil || root.Right == preNode {
			res = append(res, root.Val)
			myStack.pop()
			preNode = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return res
}

