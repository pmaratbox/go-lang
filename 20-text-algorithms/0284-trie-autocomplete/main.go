package main

import (
	"fmt"
	"sort"
	"strings"
)

type node struct {
	children map[byte]*node
	end      bool
}

func newNode() *node {
	return &node{children: make(map[byte]*node)}
}

type trie struct{ root *node }

func newTrie() *trie { return &trie{root: newNode()} }

func (t *trie) insert(w string) {
	cur := t.root
	for i := 0; i < len(w); i++ {
		c := w[i]
		if cur.children[c] == nil {
			cur.children[c] = newNode()
		}
		cur = cur.children[c]
	}
	cur.end = true
}

func (t *trie) autocomplete(prefix string) []string {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		cur = cur.children[prefix[i]]
		if cur == nil {
			return nil
		}
	}
	var res []string
	var dfs func(n *node, path string)
	dfs = func(n *node, path string) {
		if n.end {
			res = append(res, prefix+path)
		}
		keys := make([]byte, 0, len(n.children))
		for c := range n.children {
			keys = append(keys, c)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, c := range keys {
			dfs(n.children[c], path+string(c))
		}
	}
	dfs(cur, "")
	return res
}

func main() {
	t := newTrie()
	for _, w := range []string{"car", "card", "dog"} {
		t.insert(w)
	}
	fmt.Println(strings.Join(t.autocomplete("car"), " "))
}
