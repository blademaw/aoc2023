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

func buildTree(nodes []string) []*Node {
	// Incrementally build the tree with edges, if we can't find a node already,
	// create one
	nodeMap := make(map[string]*Node)
	var ghosts []*Node

	r := regexp.MustCompile(`\w{3}`)

	for _, nodeStr := range nodes {
		matches := r.FindAllString(nodeStr, -1)
		if matches[0][2] == 'A' {
			ghosts = append(ghosts, findNode(matches[0], nodeMap))
		}

		addLR(
			findNode(matches[0], nodeMap),
			findNode(matches[1], nodeMap),
			findNode(matches[2], nodeMap),
		)
	}

	return ghosts
}

// Execute a set of L/R (left-right) instructions on a node until a sink is
// reached
func executeInstructions(ghost *Node, instructions string) int {
	count, i := 0, 0

	for ghost.Value[2] != 'Z' {
		if i == len(instructions) {
			i = 0
		}

		if instructions[i] == 'L' {
			ghost = ghost.Left
		} else {
			ghost = ghost.Right
		}

		count++
		i++
	}

	return count
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(ns []int) int {
	if len(ns) < 2 {
		return ns[0]
	}

	a, b := ns[0], ns[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(ns); i++ {
		result = LCM(append([]int{result}, ns[i:]...))
	}

	return result
}

func main() {
	file := flag.String("file", "data.txt", "the .txt graph and instruction set.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	ghosts := buildTree(lines[2:])

	ghostArr := make([]int, len(ghosts))
	for i, ghostNode := range ghosts {
		ghostArr[i] = executeInstructions(ghostNode, lines[0])
	}

	fmt.Println(LCM(ghostArr))
}
