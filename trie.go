package rex

import "strings"

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{newTrieNode("")}
}

func (t *trie) add(route string) {
	route = strings.TrimPrefix(route, "/")
	node := t.root

	for _, val := range strings.Split(route, "/") {
		node = node.findOrCreate(val)
	}

	node.leaf = true
	node.route = route
}

func (t *trie) match(path string) string {
	path = strings.TrimPrefix(path, "/")
	node := t.root

	for _, val := range strings.Split(path, "/") {
		node = node.next(val)
		if node == nil {
			return ""
		}
	}

	if node.leaf {
		return "/" + node.route
	}

	return ""
}
