package bt

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
		insert(b.root, k)
	}
}

func (b *BinaryTree) Traverse(mode Mode, fn func(k Key)) {
	b.root.traverse(mode, fn)
}

func insert(parent *Node, k Key) {
	if parent == nil {
		p := NewNode(k)
		*parent = *p
	} else {
		if k.Less(parent.k) {
			if parent.l == nil {
				parent.l = NewNode(k)
			} else {
				insert(parent.l, k)
			}
		} else {
			if parent.r == nil {
				parent.r = NewNode(k)
			} else {
				insert(parent.r, k)
			}
		}
	}
}

func (b *BinaryTree) BalanceFactor() int {
	return b.root.BalanceFactor()
}
