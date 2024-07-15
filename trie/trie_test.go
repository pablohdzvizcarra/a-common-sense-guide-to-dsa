package trie

import "testing"

func TestTrieNodeCreation(t *testing.T) {
	node := TrieNode{
		children: map[rune]*TrieNode{'c': nil},
	}

	trie := NewTrie(node)

	if trie.root.children['c'] == nil {
		t.Errorf("children node in trie, must not be nil")
	}
}
