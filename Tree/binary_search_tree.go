package Tree

import "fmt"

//二叉查找树--Binary Sort Tree
/*
 * 二叉查找树定义：又称为是二叉排序树（Binary Sort Tree）或二叉搜索树。二叉排序树或者是一棵空树，或者是具有下列性质的二叉树：
 *　　　　1) 若左子树不空，则左子树上所有结点的值均小于它的根结点的值；
 *　　　　2) 若右子树不空，则右子树上所有结点的值均大于或等于它的根结点的值；
 *　　　　3) 左、右子树也分别为二叉排序树；
 *　　　　4) 没有键值相等的节点。
 * 极端情况下会退化为链表
 */

//二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

//二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64                 //值
	Times int64                 //值出现的次数
	Left  *BinarySearchTreeNode //左子树
	Right *BinarySearchTreeNode //右子树
}

//初始化一个二叉查找树
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

//添加元素--tree
func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

//node
func (node *BinarySearchTreeNode) Add(value int64) {
	// 如果插入的值比节点的值小，那么要插入到该节点的左子树中
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			node.Left.Add(value)
		}
	} else if value > node.Value { // 如果插入的值比节点的值大，那么要插入到该节点的右子树中
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			node.Right.Add(value)
		}
	} else {
		node.Times = node.Times + 1
	}
}

//找出最小值的节点
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	// 左子树为空，说明已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

//找出最大值的节点
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

//查找指定的元素
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if node.Value == value { // 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if value < node.Value { // 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	} else { // 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	}
}

//查找指定节点的父亲
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	// 如果根节点等于该值，根节点其没有父节点，返回nil
	if tree.Root.Value == value {
		return nil
	}
	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	// 外层没有值相等的判定，因为在内层已经判定完毕后返回父亲节点。
	if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		leftTree := node.Left
		if leftTree == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}

		// 左子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		rightTree := node.Right
		if rightTree == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}

		// 右子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

// 删除指定的元素
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}

	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 不存在该值，直接返回
		return
	}

	// 查找该值的父亲节点
	parent := tree.Root.FindParent(value)

	// 第一种情况，删除的是根节点，且根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除的节点有父亲节点，但没有子树

		// 如果删除的是节点是父亲的左儿子，直接将该值删除即可
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			// 删除的原来是父亲的右儿子，直接将该值删除即可
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换后二叉查找树的性质又满足了。

		// 找右子树中最小的值，一直往右子树的左边找
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 把最小的节点删掉
		tree.Delete(minNode.Value)

		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可

		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}


// 中序遍历
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 打印右子树
	node.Right.MidOrder()
}


