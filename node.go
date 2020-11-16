package main

import (
	"fmt"
)

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

// func recursiveAdd(parent *Node, derivations map[string][]string, dir string, depth int) {
// 	fmt.Printf("Depth is %d.\n", depth)
// 	if len(derivations) == 0 {
// 		fmt.Println("no more input derivations")
// 		fmt.Println()
// 		return
// 	}
// 	nixcommand := "nix show-derivation "

// 	for derivation := range derivations {
// 		command := exec.Command("/bin/bash", "-c", nixcommand+derivation, dir)
// 		out := run(command)

// 		dec := json.NewDecoder(strings.NewReader((out)))
// 		for {
// 			var derivation Derivations
// 			if err := dec.Decode(&derivation); err == io.EOF {
// 				break
// 			} else if err != nil {
// 				log.Fatal(err)
// 			}

// 			for deriv, v := range derivation {
// 				fmt.Println(deriv)
// 				node := newNode(deriv)
// 				addChild(parent, node)
// 				recursiveAdd(node, v.InputDerivations, dir, depth+1)
// 			}
// 		}
// 	}
// }

func recursiveAdd(parent *Node, derivations map[string][]string, dictionary Derivations, depth int) {
	fmt.Printf("Depth is %d.\n", depth)

	if len(derivations) == 0 {
		fmt.Println("no more input derivations")
		fmt.Println()
		return
	}

	for derivation := range derivations {
		node := newNode(derivation)
		addChild(parent, node)

		recursiveAdd(node, dictionary[derivation].InputDerivations, dictionary, depth+1)
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
