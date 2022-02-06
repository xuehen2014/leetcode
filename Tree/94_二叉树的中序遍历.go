package Tree

/*
 *  https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 */

//递归中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	inTraversal(root, &res)
	return res
}

func inTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inTraversal(root.Left, res)
	*res = append(*res, root.Val)
	inTraversal(root.Right, res)
}

//非递归中序遍历
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var myStack *myStack
	myStack = newMyStack()
	for root != nil || myStack.length() > 0 {
		for root != nil {
			myStack.push(root)
			root = root.Left
		}
		node := myStack.pop()
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}





