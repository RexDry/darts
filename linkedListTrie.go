package darts

type linkedListTrieNode struct {
	code               rune
	depth, left, right uint64
	subKey             []rune
	children           [](*linkedListTrieNode)
}

type linkedListTrie struct {
	root *linkedListTrieNode
}
