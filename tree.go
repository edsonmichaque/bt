package bt

type Tree interface {
	Inserte(Key)
	Traverse(Mode, func(Key))
}
