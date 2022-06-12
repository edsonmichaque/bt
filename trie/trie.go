package trie

type Node struct {
	children [26]*Node
	terminal bool
}

func NewNode() *Node {
	return &Node{
		terminal: false,
	}
}

func (n *Node) Insert(text string) {
	ptr := n

	for i, c := range text {
		idx := index(c)
		if ptr.children[idx] == nil {
			ptr.children[idx] = NewNode()
		}

		ptr = ptr.children[idx]
		if len(text) == 1 {
			ptr.terminal = true
			return
		} else {
			ptr.Insert(text[i+1:])
		}
	}
}

func (n *Node) Search(ss string) bool {
	for i, c := range ss {
		idx := index(c)

		child := n.children[idx]
		if child == nil {
			return false
		}

		if child.terminal && i == len(ss)-1 {
			return true
		} else {
			return child.Search(ss[i+1:])
		}
	}

	return false
}

func index(r rune) int {
	return int(r-'a') % 26
}
