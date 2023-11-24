package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and hahs a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // Letter ASCII value - 97 (ASCII value a)
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that word
// is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // Letter ASCII value - 97 (ASCII value a)
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	//Check if the lastNode is true or not
	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()
	//fmt.Println(myTrie)      // > &{0xc000094000}
	//fmt.Println(myTrie.root) // > &{[<nil> <nil> ...] false}

	toAdd := []string{
		"aragorn",
		"aragon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.root) // > &{[0xc0001320e0 <nil> <nil> <nil> 0xc0001327e0
	// <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>
	// 0xc000132d20 <nil> <nil> <nil> <nil> <nil> <nil> <nil>
	// <nil> <nil> <nil> <nil>] false
	fmt.Println(myTrie.Search("aragorn"))
}
