# Clone a Graph

> Striver SDE Sheet 152 · Day 23 Graph — one original→copy map serves as visited set and answer book; registering before recursing is what tames cycles. Go, stdlib only, tested.

## Problem

Deep-copy an undirected graph given one node: every node duplicated, every adjacency preserved among the duplicates, no reference back to an original. The sheet's parenthetical — "not as easy as it looks" — is about cycles.

## Approach

DFS with a map from original to clone. The map answers two questions at once: "have I been here?" and "what do I wire this edge to?" — one structure instead of a visited set plus a lookup table. The ordering subtlety that breaks naive versions: the clone must be **registered in the map before recursing** into neighbors; otherwise a cycle (even the minimal two-node one) re-enters the node before its clone exists and recurses forever. The test clones a 4-cycle and asserts the cloned cycle closes onto the *same* clone objects and never touches an original.

**Complexity:** O(V+E) time, O(V) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
