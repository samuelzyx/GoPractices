package main

import "fmt"

var count int

//Node represents the components of a
//binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will add a node to the tree
func (n *Node) Insert(k int) {
	//Recursive
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search will take in a key value
//and RETURN true if there is a
//node with that value
func (n *Node) Search(k int) bool {
	count++

	//if node is nil is end of tree
	//is not exist match
	if n == nil {
		return false
	}

	//Recursive
	if n.Key < k {
		//move left
		return n.Right.Search(k)
	} else if n.Key > k {
		//move right
		return n.Left.Search(k)
	}

	//is exist match
	return true
}

func main() {
	tree := &Node{Key: 100} //Node 1
	tree.Insert(150)        //Node 2
	tree.Insert(7)          //Node 2
	tree.Insert(50)         //Node 3
	tree.Insert(60)         //Node 4
	tree.Insert(97)         //Node 5
	tree.Insert(6)          //Node 3
	tree.Insert(500)        //Node 3
	tree.Insert(255)        //Node 4
	tree.Insert(1000)       //Nodo 4

	// Node1	                					100
	// Node2					7											150
	// Node3			6				50						nil							500
	// Node4		nil		nil		nil		60				nil		nil				255				100
	// Node5							nil		97								nil		nil		nil		nil
	// Node6								nil		nil

	fmt.Println(tree.Search(97))
	fmt.Println(count)

	//Complexity: O(log n)
}
