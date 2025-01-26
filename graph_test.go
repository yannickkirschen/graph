package graph_test

import (
	"testing"

	"github.com/yannickkirschen/graph"
)

/*
       c a,-4-,bc
    1 - 2 - 3 - 5 - 6
   a b a b a b b a a b
*/

func TestFind(t *testing.T) {
	one := graph.NewNode[int, string](1)
	two := graph.NewNode[int, string](2)
	three := graph.NewNode[int, string](3)
	four := graph.NewNode[int, string](4)
	five := graph.NewNode[int, string](5)
	six := graph.NewNode[int, string](6)

	one.ConnectBi("a", "b")
	two.ConnectBi("a", "b")
	two.ConnectBi("a", "c")
	three.ConnectBi("a", "b")
	four.ConnectBi("a", "b")
	five.ConnectBi("a", "b")
	five.ConnectBi("a", "c")
	six.ConnectBi("a", "b")

	graph := graph.NewGraph[int, string]()
	graph.AddNode(one)
	graph.AddNode(two)
	graph.AddNode(three)
	graph.AddNode(four)
	graph.AddNode(five)
	graph.AddNode(six)

	graph.ConnectRefBi(1, "b", 2, "a")
	graph.ConnectRefBi(2, "b", 3, "a")
	graph.ConnectRefBi(2, "c", 4, "a")
	graph.ConnectRefBi(3, "b", 5, "b")
	graph.ConnectRefBi(4, "b", 5, "c")
	graph.ConnectRefBi(5, "a", 6, "a")

	paths := graph.FindRef(1, "b", 4, "b")
	if len(paths) != 1 {
		t.Fatalf("expected 1 path, but got %d: %v", len(paths), paths)
	}

	path := paths[0]
	if len(path) != 3 {
		t.Fatalf("path has wrong length: expected 3, but got %d: %v", len(path), path)
	}

	if path[0].Id() != 1 || path[1].Id() != 2 || path[2].Id() != 4 {
		t.Fatalf("expected path to be 1 -> 2 -> 4 but got %v", path)
	}
}
