package edge

import (
	"math"

	"github.com/edsonmichaque/pungue"
)

var (
	InfinityInt = Int{
		Value: math.MaxInt,
	}
)

type Int struct {
	Value int
}

func (v Int) Add(other pungue.Edge) pungue.Edge {
	return Int{
		Value: v.Value + other.(Int).Value,
	}
}

func (v Int) Eq(other pungue.Edge) bool {
	return v.Value == other.(Int).Value
}

func (v Int) Diff(other pungue.Edge) bool {
	return v.Value != other.(Int).Value
}

func (v Int) Gt(other pungue.Edge) bool {
	return v.Value > other.(Int).Value
}

func (v Int) Lt(other pungue.Edge) bool {
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
