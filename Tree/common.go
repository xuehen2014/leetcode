package Tree


type myStack struct {
	content []*TreeNode
	top     int
}

func newMyStack() *myStack {
	return &myStack{
		content: []*TreeNode{},
		top:     -1,
	}
}

func (m *myStack) push(node *TreeNode) {
	m.content = append(m.content, node)
	m.top++
}

func (m *myStack) pop() *TreeNode {
	val := m.content[m.top]
	m.content = m.content[:m.top]
	m.top--
	return val
}

func (m *myStack) peek() *TreeNode {
	return m.content[m.top]
}

func (m *myStack) length() int {
	return len(m.content)
}
