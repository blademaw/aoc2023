package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"regexp"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Find a node with a name->node map; otherwise create it and return a pointer
// to it
func findNode(name string, nmap map[string]*Node) *Node {
	if node, ok := nmap[name]; ok {
		return node
	} else {
		node := Node{Value: name, Left: nil, Right: nil}
		nmap[name] = &node
		return &node
	}
}

// Add left and right pointers to a node
func addLR(node *Node, left *Node, right *Node) {
	node.Left = left
	node.Right = right
}

func buildTree(nodes []string) Tree {
	// Incrementally build the tree with edges, if we can't find a node already,
	// create one
	nodeMap := make(map[string]*Node)

	r := regexp.MustCompile(`\w{3}`)

	for _, nodeStr := range nodes {
		matches := r.FindAllString(nodeStr, -1)

		addLR(
			findNode(matches[0], nodeMap),
			findNode(matches[1], nodeMap),
			findNode(matches[2], nodeMap),
		)
	}

	return Tree{Root: nodeMap["AAA"]}
}

// Execute a set of L/R (left-right) instructions on a tree
func executeInstructions(tree Tree, instructions string) int {
	count, i, node := 0, 0, tree.Root

	for node.Value != "ZZZ" {
		if i == len(instructions) {
			i = 0
		}

		if instructions[i] == 'L' {
			node = node.Left
		} else {
			node = node.Right
		}

		count++
		i++
	}

	return count
}

func main() {
	file := flag.String("file", "data.txt", "the .txt graph and instruction set.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	tree := buildTree(lines[2:])
	fmt.Println(executeInstructions(tree, lines[0]))
}
