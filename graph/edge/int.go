package edge

import (
	"math"

	"github.com/edsonmichaque/pungue/graph"
)

var (
	InfinityInt = Int{
		Value: math.MaxInt,
	}
)

type Int struct {
	Value int
}

func (v Int) Add(other graph.Edge) graph.Edge {
	return Int{
		Value: v.Value + other.(Int).Value,
	}
}

func (v Int) Eq(other graph.Edge) bool {
	return v.Value == other.(Int).Value
}

func (v Int) Diff(other graph.Edge) bool {
	return v.Value != other.(Int).Value
}

func (v Int) Gt(other graph.Edge) bool {
	return v.Value > other.(Int).Value
}

func (v Int) Lt(other graph.Edge) bool {
	return v.Value < other.(Int).Value
}

func (v Int) Inf() bool {
	return v.Value == math.MaxInt
}

func (v Int) Zero() bool {
	return v.Value == 0
}

func NewInt(v int) Int {
	return Int{
		Value: v,
	}
}
