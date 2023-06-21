package main

type Trie struct {
	root *trieNode
}

type trieNode struct {
	end  bool
	next map[byte]*trieNode
}

func ConstructorTrie() Trie {
	return Trie{
		root: &trieNode{
			next: map[byte]*trieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i, c := range []byte(word) {
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			node := &trieNode{
				next: map[byte]*trieNode{},
			}
			cur.next[c] = node
			cur = node
		}
		if i == len(word)-1 {
			cur.end = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range []byte(word) {
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			return false
		}
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range []byte(prefix) {
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			return false
		}
	}
	return true
}
