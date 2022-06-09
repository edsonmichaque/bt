package libtree

import "fmt"

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
