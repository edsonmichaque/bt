package pungue

import (
	"errors"
	"math"
)

type Graph struct {
	vertices        map[int]Vertex
	adjacencyMatrix [][]int
}

type Vertex interface {
	Equal(Vertex) bool
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v Vertex) error {
	for _, e := range g.vertices {
		if v.Equal(e) {
			return errors.New("graph: cannot add vertex")
		}
	}

	g.vertices[len(g.vertices)] = v

	aux := newAdjacencyMatrix(len(g.adjacencyMatrix) + 1)
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
		if v.Equal(e) {
			pos = i
			break
		}
	}

	if pos == math.MaxInt {
		return errors.New("graph: invalid vertex")
	}

	arr := newAdjacencyMatrix(len(g.adjacencyMatrix) - 1)

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

func newAdjacencyMatrix(l int) [][]int {
	arr := make([][]int, l)
	for i := 0; i < l; i++ {
		arr = append(arr, repeat(0, l))
	}

	return arr[:]
}

func repeat(v, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, v)
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
