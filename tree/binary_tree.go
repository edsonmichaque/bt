package tree

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) Insert(k Key) {
	if b.root == nil {
		b.root = NewNode(k)
	} else {
		b.root.insert(k)
	}
}

func (b *BinaryTree) Traverse(mode Mode, fn func(k Key)) {
	b.root.traverse(mode, fn)
}

func (parent *Node) insert(k Key) {
	if parent == nil {
		p := NewNode(k)
		*parent = *p
	} else {
		if k.Lt(parent.k) {
			if parent.l == nil {
				parent.l = NewNode(k)
				parent.l.p = parent
			} else {
				parent.l.insert(k)
			}
		} else {
			if parent.r == nil {
				parent.r = NewNode(k)
				parent.r.p = parent
			} else {
				parent.r.insert(k)

			}
		}
	}
}

func (b *BinaryTree) BalanceFactor() int {
	return b.root.BalanceFactor()
}
