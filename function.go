package main

import (
	"bufio"
	"io"
	"strings"
)

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
}

func parseInput(input string) (string, string) {
	a := strings.Split(input, ")")
	return a[0], a[1]
}

func buildMap(r io.Reader) *Node {
	nodeMap := make(map[string]*Node)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var (
			found               bool
			leftNode, rightNode *Node
		)
		s1, s2 := parseInput(scanner.Text())
		if leftNode, found = nodeMap[s1]; !found {
			leftNode = new(Node)
			nodeMap[s1] = leftNode
			leftNode.Name = s1
		}
		if rightNode, found = nodeMap[s2]; !found {
			rightNode = new(Node)
			rightNode.Name = s2
			nodeMap[s2] = rightNode
		}
		leftNode.Children = append(leftNode.Children, rightNode)
		if rightNode.Parent != nil {
			panic("More than one parent for node: " + rightNode.Name)
		}
		rightNode.Parent = leftNode
	}
	return nodeMap["COM"]
}

func getAllDepths(root *Node) int {
	depth := getDepth(root)
	for _, p := range root.Children {
		depth += getAllDepths(p)
	}
	return depth
}

func getDepth(root *Node) int {
	depth := 0
	for _, p := range root.Children {
		depth++
		depth += getDepth(p)
	}
	return depth
}

func findNode(root *Node, target string) *Node {
	if root.Name == target {
		return root
	}
	for _, p := range root.Children {
		if x := findNode(p, target); x != nil {
			return x
		}
	}
	return nil
}

func findDistToNode(startNode *Node, targetNode *Node) int {
	if d := findDistToChild(startNode, targetNode); d > 0 {
		return d
	}
	if startNode.Parent != nil {
		return findDistToNode(startNode.Parent, targetNode) + 1
	}
	return 0
}

func findDistToChild(startNode *Node, targetNode *Node) int {
	if startNode != targetNode {
		for _, p := range startNode.Children {
			if p == targetNode {
				return 1
			}
			if d := findDistToChild(p, targetNode); d > 0 {
				return 1 + d
			}
		}
	}
	return 0
}
