package main

import "fmt"

type Node struct {
	derivation string
	children   []*Node
}

func newNode(derivation string) *Node {
	var node Node
	node.derivation = derivation
	return &node
}

func addChild(parent *Node, child *Node) {
	parent.children = append(parent.children, child)
}

func numChildren(parent *Node) int {
	return len(parent.children)
}

func recursiveAdd(parent *Node, derivations map[string][]string) {
	for derivation := range derivations {
		fmt.Println(derivation)
	}
}

// Testing funcs
// func main() {
// 	s := make([]string, 0)
// 	s = append(s, "hello")
// 	s = append(s, "yo")
// 	fmt.Println(s)

// 	node1 := newNode("node1")
// 	node2 := newNode("node2")
// 	node3 := newNode("node3")

// 	addChild(node2, node3)
// 	addChild(node1, node2)
// 	for index, element := range node1.children {
// 		fmt.Println("Child of Node 1")
// 		if element.derivation != "" {
// 			fmt.Println(index, "--", element.derivation)
// 		}
// 		fmt.Println("Child of Node2")
// 		for i, e := range element.children {
// 			if e.derivation != "" {
// 				fmt.Println(i, "--", e.derivation)
// 			}
// 		}
// 	}

// 	//fmt.Println(numChildren(&node1))
// 	//fmt.Println(node1.children[0].children[0].derivation)
// }
