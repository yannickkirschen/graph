package graph

import (
	"fmt"
	"slices"
)

type Node[I comparable, P comparable] struct {
	id          I
	connections map[P][]P
}

func NewNode[I comparable, P comparable](id I) *Node[I, P] {
	return &Node[I, P]{id, map[P][]P{}}
}

func (node *Node[I, P]) Id() I {
	return node.id
}

func (node *Node[I, P]) String() string {
	return fmt.Sprintf("Node<%v>", node.id)
}

func (node *Node[I, P]) Connect(from, to P) {
	connection, ok := node.connections[from]
	if !ok {
		connection = []P{}
	}

	node.connections[from] = append(connection, to)
}

func (node *Node[I, P]) ConnectBi(from, to P) {
	node.Connect(from, to)
	node.Connect(to, from)
}

func (node *Node[I, P]) Next(port P) []P {
	return node.connections[port]
}

type vertex[I comparable, P comparable] struct {
	id   I
	port P
}

type Graph[I comparable, P comparable] struct {
	nodes       map[I]*Node[I, P]
	connections map[vertex[I, P]]vertex[I, P]
}

func NewGraph[I comparable, P comparable]() *Graph[I, P] {
	return &Graph[I, P]{map[I]*Node[I, P]{}, map[vertex[I, P]]vertex[I, P]{}}
}

func (graph *Graph[I, P]) AddNode(node *Node[I, P]) {
	graph.nodes[node.id] = node
}

func (graph *Graph[I, P]) Connect(from *Node[I, P], fromPort P, to *Node[I, P], toPort P) {
	graph.ConnectRef(from.id, fromPort, to.id, toPort)
}

func (graph *Graph[I, P]) ConnectBi(from *Node[I, P], fromPort P, to *Node[I, P], toPort P) {
	graph.ConnectRef(from.id, fromPort, to.id, toPort)
	graph.ConnectRef(to.id, toPort, from.id, fromPort)
}

func (graph *Graph[I, P]) ConnectRef(from I, fromPort P, to I, toPort P) {
	graph.connections[vertex[I, P]{from, fromPort}] = vertex[I, P]{to, toPort}
}

func (graph *Graph[I, P]) ConnectRefBi(from I, fromPort P, to I, toPort P) {
	graph.ConnectRef(from, fromPort, to, toPort)
	graph.ConnectRef(to, toPort, from, fromPort)
}

func (graph *Graph[I, P]) Find(from *Node[I, P], fromPort P, to *Node[I, P], toPort P) [][]*Node[I, P] {
	data := &dFSData[I, P]{[]*Node[I, P]{}, []*Node[I, P]{}, [][]*Node[I, P]{}}
	graph.dfsFind(data, from, fromPort, to, toPort)
	return data.Paths
}

func (graph *Graph[I, P]) FindRef(from I, fromPort P, to I, toPort P) [][]*Node[I, P] {
	start, ok := graph.nodes[from]
	if !ok {
		panic("from not found")
	}
	end, ok := graph.nodes[to]
	if !ok {
		panic("end not found")
	}

	return graph.Find(start, fromPort, end, toPort)
}

func (graph *Graph[I, P]) dfsFind(data *dFSData[I, P], from *Node[I, P], fromPort P, to *Node[I, P], toPort P) {
	if slices.Contains(data.Visited, from) {
		return
	}

	data.Visited = append(data.Visited, from)
	data.CurrentPath = append(data.CurrentPath, from)

	if from.id == to.id && fromPort == toPort {
		var currentPath []*Node[I, P]
		currentPath = append(currentPath, data.CurrentPath...)

		data.Paths = append(data.Paths, currentPath)
		data.Visited = data.Visited[:len(data.Visited)-1]
		data.CurrentPath = data.CurrentPath[:len(data.CurrentPath)-1]
		return
	}

	nextVertex, ok := graph.connections[vertex[I, P]{from.id, fromPort}]
	if !ok {
		data.CurrentPath = data.CurrentPath[:len(data.CurrentPath)-1]
		return
	}

	next := graph.nodes[nextVertex.id]
	for _, port := range next.Next(nextVertex.port) {
		graph.dfsFind(data, next, port, to, toPort)
	}

	data.CurrentPath = data.CurrentPath[:len(data.CurrentPath)-1]
	idx := slices.Index(data.Visited, from)
	data.Visited = slices.Delete(data.Visited, idx, idx+1)
}

type dFSData[I comparable, P comparable] struct {
	Visited     []*Node[I, P]
	CurrentPath []*Node[I, P]
	Paths       [][]*Node[I, P]
}
