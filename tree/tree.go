package tree

type Tree interface {
	Inserte(Key)
	Traverse(TraversalMode, func(Key))
}
