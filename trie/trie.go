package trie

import (
	"fmt"
	"unicode/utf8"
)

type Trie struct {
	IsStr bool
	R     rune
	Next  map[rune]*Trie
}

func New() *Trie {
	return &Trie{IsStr: false, Next: nil}
}
func (t *Trie) Search(str string) bool {

}

func (t *Trie) Insert(str string) {
	root := t
	b := []byte(str)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == utf8.RuneError {
			r = rune(b[0])
		}
		if root.Next == nil {
			root.Next = make(map[rune]*Trie)
		}
		if root.Next[r] == nil {
			root.Next[r] = New()
			root.Next[r].R = r
		}
		root = root.Next[r]
		b = b[size:]
	}
	root.IsStr = true
}

func main() {
	fmt.Println("vim-go")
}
