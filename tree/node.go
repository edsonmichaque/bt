package tree

type Node struct {
	k Key
	p *Node
	l *Node
	r *Node
}

type Key interface {
	Lt(Key) bool
}

func NewNode(v Key) *Node {
	return &Node{k: v}
}

func (n *Node) Height() int {
	if n != nil {
		leftHeight := 0
		if n.l != nil {
			leftHeight = n.l.Height() + 1
		}

		rightHeight := 0
		if n.r != nil {
			rightHeight = n.r.Height() + 1
		}

		if leftHeight > rightHeight {
			return leftHeight
		}

		return rightHeight
	}

	return 0
}

func (n *Node) BalanceFactor() int {
	return n.l.Height() - n.r.Height()
}

func (n *Node) traverse(mode TraversalMode, fn func(k Key)) {
	if n != nil {
		if mode == Postorder {
			n.l.traverse(mode, fn)
			n.r.traverse(mode, fn)
			fn(n.k)
		} else if mode == Preorder {
			fn(n.k)
			n.l.traverse(mode, fn)
			n.r.traverse(mode, fn)
		} else {
			n.l.traverse(mode, fn)
			fn(n.k)
			n.r.traverse(mode, fn)
		}
	}
}
