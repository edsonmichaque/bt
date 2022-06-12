package pungue

type Edge interface {
	Add(Edge) Edge
	Lt(Edge) bool
	Gt(Edge) bool
	Eq(Edge) bool
	Inf() bool
	Zero() bool
}
