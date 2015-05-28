package rex

import "regexp"

type nodeType uint8

const (
	staticNode nodeType = iota
	paramNode
)

var (
	paramNodeRx = regexp.MustCompile(`^:.*`)
)

type trieNode struct {
	val string
	typ nodeType

	leaf  bool
	route string

	edges []*trieNode
}

func (node *trieNode) findOrCreate(val string) *trieNode {
	for _, edge := range node.edges {
		if edge.equal(val) {
			return edge
		}
	}

	edge := newTrieNode(val)
	node.edges = append(node.edges, edge)

	return edge
}

func (node *trieNode) equal(val string) bool {
	if node.typ == paramNode {
		if paramNodeRx.MatchString(val) {
			return true
		} else {
			return false
		}
	}

	return node.val == val
}

func (node *trieNode) next(val string) *trieNode {
	for _, edge := range node.edges {
		if edge.typ == paramNode {
			return edge
		} else {
			if edge.val == val {
				return edge
			}
		}
	}

	return nil
}

func newTrieNode(val string) *trieNode {
	typ := staticNode
	if paramNodeRx.MatchString(val) {
		typ = paramNode
		val = ""
	}

	return &trieNode{
		val:   val,
		typ:   typ,
		edges: []*trieNode{},
	}
}
