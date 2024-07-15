package trie

type TrieNode struct {
	children map[rune]*TrieNode
}

type Trie struct {
	root TrieNode
}

func NewTrie(node TrieNode) *Trie {
	return &Trie{
		root: node,
	}
}

func (t *Trie) search(word string) *TrieNode {
	currentNode := &t.root

	for _, char := range word {
		if _, ok := currentNode.children[char]; ok {
			currentNode = currentNode.children[char]
		} else {
			return nil
		}
	}

	return currentNode
}
