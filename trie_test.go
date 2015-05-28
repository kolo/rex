package rex

import (
	"fmt"
	"testing"
)

var testRoutes = []string{
	"/login",
	"/user/:id",
	"/user/:id/photos",
	"/user/:user_id/photos/:id",
}

var trieTests = []struct {
	path  string
	route string
}{
	{"/login", "/login"},
	{"/user/7", "/user/:id"},
	{"/user/7/photos", "/user/:id/photos"},
	{"/user/7/photos/3", "/user/:user_id/photos/:id"},
}

func Test_trieAdd(t *testing.T) {
	routes := newTrie()
	for _, r := range testRoutes {
		routes.add(r)
	}

	printTrieNode("->", routes.root)

	for _, tt := range trieTests {
		route := routes.match(tt.path)
		if route != tt.route {
			t.Fatalf("Expected: match(%q)=%q\nGot: match(%q)=%q\n", tt.path, tt.route, tt.path, route)
		}
	}
}

func printTrieNode(prefix string, node *trieNode) {
	fmt.Printf("%s/%s(%d)\n", prefix, node.val, node.typ)
	prefix = "-" + prefix
	for _, edge := range node.edges {
		printTrieNode(prefix, edge)
	}
}
