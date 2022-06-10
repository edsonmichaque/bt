package libtree

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

	aux := make([][]int, len(g.adjacencyMatrix)-1)
	for i := 0; i < len(g.adjacencyMatrix)-1; i++ {
		aux = append(aux, repeat(0, len(g.adjacencyMatrix)-2))
	}

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

	aux := make([][]int, len(g.adjacencyMatrix)-1)
	for i := 0; i < len(g.adjacencyMatrix)-1; i++ {
		aux = append(aux, repeat(0, len(g.adjacencyMatrix)-2))
	}

	return nil
}

func repeat(v, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, v)
	}

	return arr
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
