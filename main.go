package main

import "fmt"

// Node is an undirected graph node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// Clone deep-copies a connected graph. The original→copy map is both the
// visited set and the answer book: a node already mapped is returned, not
// re-cloned — which is exactly what makes cycles terminate.
func Clone(n *Node) *Node {
	copies := map[*Node]*Node{}
	var dfs func(cur *Node) *Node
	dfs = func(cur *Node) *Node {
		if cur == nil {
			return nil
		}
		if c, ok := copies[cur]; ok {
			return c
		}
		c := &Node{Val: cur.Val}
		copies[cur] = c // register BEFORE recursing, or cycles loop forever
		for _, nb := range cur.Neighbors {
			c.Neighbors = append(c.Neighbors, dfs(nb))
		}
		return c
	}
	return dfs(n)
}

func main() {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	a.Neighbors = []*Node{b}
	b.Neighbors = []*Node{a}
	c := Clone(a)
	fmt.Println(c.Val, c.Neighbors[0].Val, c.Neighbors[0].Neighbors[0] == c)
}
