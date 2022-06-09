package bt

import (
	"fmt"
)

type Tree interface {
	Inserte(Key)
	Traverse(Mode, func(Key))
}

func main() {
	bt := NewBinaryTree()
	bt.Insert(Int(4))
	bt.Insert(Int(2))
	bt.Insert(Int(1))
	bt.Insert(Int(3))
	bt.Insert(Int(5))

	bt.Traverse(Inorder, func(v Key) {
		fmt.Println(v)
	})

	balance := bt.BalanceFactor()
	fmt.Println("Balance: ", balance)
}

type TreeNode struct {
	key        Key
	leftchild  *TreeNode
	rightchild *TreeNode
}

type Key interface {
	Less(Key) bool
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func NewTreeNode(v Key) *TreeNode {
	return &TreeNode{key: v}
}

type BinaryTree struct {
	root *TreeNode
}

func (b *BinaryTree) Insert(k Key) {
	if b.root == nil {
		b.root = NewTreeNode(k)
	} else {
		insert(b.root, k)
	}
}

type Mode int

const (
	Inorder Mode = iota
	Preorder
	Postorder
)

func (b *BinaryTree) Traverse(mode Mode, fn func(k Key)) {
	traverse(b.root, mode, fn)
}

func insert(parent *TreeNode, k Key) {
	if parent == nil {
		p := NewTreeNode(k)
		*parent = *p
	} else {
		if k.Less(parent.key) {
			if parent.leftchild == nil {
				parent.leftchild = NewTreeNode(k)
			} else {
				insert(parent.leftchild, k)
			}
		} else {
			if parent.rightchild == nil {
				parent.rightchild = NewTreeNode(k)
			} else {
				insert(parent.rightchild, k)
			}
		}
	}
}

func traverse(n *TreeNode, mode Mode, fn func(k Key)) {
	if n != nil {
		if mode == Postorder {
			traverse(n.leftchild, mode, fn)
			traverse(n.rightchild, mode, fn)
			fn(n.key)
		} else if mode == Preorder {
			fn(n.key)
			traverse(n.leftchild, mode, fn)
			traverse(n.rightchild, mode, fn)
		} else {
			traverse(n.leftchild, mode, fn)
			fn(n.key)
			traverse(n.rightchild, mode, fn)
		}
	}
}

type Int int

func (i Int) Less(o Key) bool {
	if v, ok := o.(Int); ok {
		return int(i) < int(v)
	}

	return false
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}

func balanceFactor(n *TreeNode) int {
	return height(n.leftchild) - height(n.rightchild)
}

func (b *BinaryTree) BalanceFactor() int {
	return balanceFactor(b.root)
}

func height(n *TreeNode) int {
	if n != nil {
		leftHeight := 0
		if n.leftchild != nil {
			leftHeight = height(n.leftchild) + 1
		}

		rightHeight := 0
		if n.rightchild != nil {
			rightHeight = height(n.rightchild) + 1
		}

		if leftHeight > rightHeight {
			return leftHeight
		}

		return rightHeight
	}

	return 0
}
