# Graph

[![Lint commit message](https://github.com/yannickkirschen/graph/actions/workflows/commit-lint.yml/badge.svg)](https://github.com/yannickkirschen/graph/actions/workflows/commit-lint.yml)
[![Go](https://github.com/yannickkirschen/graph/actions/workflows/go.yml/badge.svg)](https://github.com/yannickkirschen/graph/actions/workflows/go.yml)
[![GitHub release](https://img.shields.io/github/release/yannickkirschen/graph.svg)](https://github.com/yannickkirschen/graph/releases/)

A simple graph library written in Go.

Graphs are represented by an adjacency list combined with ports for each node.
Paths can be found by using recursive depth first search.

## Caution when using references

When using reference types as type parameter, always use the same pointer to
reference the same logical object! As the type is being used as key in a map and
compared for equality, this is very important.

## Usage

Consider the following simple graph:

![Image of a graph](docs/graph.png "Example graph")

### Define the nodes

```go
one := graph.NewNode[int, string](1)
two := graph.NewNode[int, string](2)
three := graph.NewNode[int, string](3)
four := graph.NewNode[int, string](4)
five := graph.NewNode[int, string](5)
six := graph.NewNode[int, string](6)
```

### Define inner port connections

```go
one.ConnectBi("a", "b")
two.ConnectBi("a", "b")
two.ConnectBi("a", "c")
three.ConnectBi("a", "b")
four.ConnectBi("a", "b")
five.ConnectBi("a", "b")
five.ConnectBi("a", "c")
six.ConnectBi("a", "b")
```

### Initialize the graph

```go
graph := graph.NewGraph[int, string]()
graph.AddNode(one)
graph.AddNode(two)
graph.AddNode(three)
graph.AddNode(four)
graph.AddNode(five)
graph.AddNode(six)
```

### Define outer node connections

```go
graph.ConnectRefBi(1, "b", 2, "a")
graph.ConnectRefBi(2, "b", 3, "a")
graph.ConnectRefBi(2, "c", 4, "a")
graph.ConnectRefBi(3, "b", 5, "b")
graph.ConnectRefBi(4, "b", 5, "c")
graph.ConnectRefBi(5, "a", 6, "a")
```

#### Find all paths

```go
paths := graph.FindRef(1, "b", 4, "b")
```
