package trie

import "testing"

func TestTrie(t *testing.T) {
	v := New()
	v.Insert("hello")
	t.Fatalf("%#v", v)

}
