package graph

type Edge interface {
	Add(Edge) Edge
	Gt(Edge) bool
	IsInfinity() bool
	IsZero() bool
}
