package tree

func (node *Node) TraverseInorder() {
	if node == nil {
		return
	}
	node.Left.TraverseInorder()
	println(node.Value)
	node.Right.TraverseInorder()
}

func TraverseInorder(root *Node) {
	if root == nil {
		return
	}
	TraverseInorder(root.Left)
	println(root.Value)
	TraverseInorder(root.Right)
}
