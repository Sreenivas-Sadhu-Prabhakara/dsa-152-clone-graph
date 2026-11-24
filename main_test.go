package main

import "testing"

func TestClone(t *testing.T) {
	// Square: 1-2-3-4-1.
	n1, n2, n3, n4 := &Node{Val: 1}, &Node{Val: 2}, &Node{Val: 3}, &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	c1 := Clone(n1)
	if c1 == n1 {
		t.Fatal("clone must be a new node")
	}
	c2 := c1.Neighbors[0]
	c4 := c1.Neighbors[1]
	if c2.Val != 2 || c4.Val != 4 {
		t.Fatalf("neighbor values wrong: %d %d", c2.Val, c4.Val)
	}
	// The cycle must close onto the SAME clone objects, not fresh ones.
	if c2.Neighbors[0] != c1 {
		t.Error("cycle back to the cloned 1 broken")
	}
	c3 := c2.Neighbors[1]
	if c3.Val != 3 || c3.Neighbors[1] != c4 {
		t.Error("far side of the square wrong")
	}
	// No original node may appear anywhere in the clone.
	seen := map[*Node]bool{}
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil || seen[n] {
			return
		}
		seen[n] = true
		if n == n1 || n == n2 || n == n3 || n == n4 {
			t.Fatal("clone references an original node")
		}
		for _, nb := range n.Neighbors {
			walk(nb)
		}
	}
	walk(c1)

	if Clone(nil) != nil {
		t.Error("nil clones to nil")
	}
}
