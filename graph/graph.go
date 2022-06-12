package graph

import (
	"errors"
	"math"
)

type Edge interface {
	Add(Edge) Edge
	Gt(Edge) bool
	Inf() bool
	Zero() bool
}

type Graph struct {
	vertices        map[int]Vertex
	adjacencyMatrix [][]int
}

type Vertex interface {
	Eq(Vertex) bool
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v Vertex) error {
	for _, e := range g.vertices {
		if v.Eq(e) {
			return errors.New("graph: cannot add vertex")
		}
	}

	g.vertices[len(g.vertices)] = v

	aux := newMatrix(len(g.adjacencyMatrix) + 1)
	for i, l := range g.adjacencyMatrix {
		for j, e := range l {
			aux[i][j] = e
		}
	}

	g.adjacencyMatrix = aux

	return nil
}

func (g *Graph) RemoveVertex(v Vertex) error {
	pos := math.MaxInt
	for i, e := range g.vertices {
		if v.Eq(e) {
			pos = i
			break
		}
	}

	if pos == math.MaxInt {
		return errors.New("graph: invalid vertex")
	}

	arr := newMatrix(len(g.adjacencyMatrix) - 1)

	for i, l := range g.adjacencyMatrix {
		if i == pos {
			continue
		}

		for j, e := range l {
			if j == pos {
				continue
			} else if i >= pos || j >= pos {
				arr[i-1][j-1] = e
			} else {
				arr[i][j] = e
			}
		}
	}

	g.adjacencyMatrix = arr[:]

	return nil
}

func newMatrix(l int) [][]int {
	arr := make([][]int, l)
	for i := 0; i < l; i++ {
		arr = append(arr, newSlice(0, l))
	}

	return arr[:]
}

func newSlice(value, size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr = append(arr, value)
	}

	return arr[:]
}

func (g *Graph) AddEdge(v, w int) error {
	if v >= len(g.vertices) || w >= len(g.vertices) {
		return errors.New("graph: invalid edges")
	}

	g.adjacencyMatrix[v][w] = 1
	return nil
}

func (g *Graph) RemoveEdge(v, w int) error {
	if v >= len(g.vertices) || w >= len(g.vertices) {
		return errors.New("graph: invalid edges")
	}

	g.adjacencyMatrix[v][w] = 0
	return nil
}

func Dijkstra(g [][]Edge, src int) [][]Edge {
	paths := make([][]Edge, len(g))
	for i, l := range g {
		paths[i] = make([]Edge, len(l))
		for j := range g {
			paths[i][j] = Inf{}
		}
	}

	visited := map[int]bool{
		src: true,
	}

	for i, l := range g {
		if _, ok := visited[i]; ok {
			for j, c := range l {
				if c != nil && !c.Zero() {
					if paths[src][j].Inf() && paths[src][i].Inf() {
						paths[src][j] = c
					} else if paths[src][j].Gt(paths[src][i].Add(c)) {
						paths[src][j] = paths[src][i].Add(c)
					}

					visited[j] = true
				}
			}
		}
	}

	return paths
}

type Inf struct {
	Value int
}

func (v Inf) Add(other Edge) Edge {
	return Inf{}
}

func (v Inf) Eq(other Edge) bool {
	return false
}

func (v Inf) Diff(other Edge) bool {
	return false
}

func (v Inf) Gt(other Edge) bool {
	return false
}

func (v Inf) Lt(other Edge) bool {
	return false
}

func (v Inf) Inf() bool {
	return true
}

func (v Inf) Zero() bool {
	return false
}

func (v Inf) String() string {
	return "+Inf"
}
