package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Node struct {
	Id        string
	Connected []*Node

	isBig   bool
	isEnd   bool
	isStart bool
}

func NewNode(id string) *Node {
	return &Node{
		Id:        id,
		Connected: []*Node{},
		isBig:     strings.ToUpper(id) == id,
		isEnd:     id == "end",
		isStart:   id == "start",
	}
}

func (node *Node) IsBig() bool {
	return node.isBig
}

func (node *Node) IsEnd() bool {
	return node.isEnd
}

func (node *Node) IsStart() bool {
	return node.isStart
}

type Today struct {
	nodes map[string]*Node
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadDelimitedFile(input, "-")
	if err != nil {
		return err
	}

	d.nodes = make(map[string]*Node)
	for _, row := range in {
		var node *Node
		var ok bool
		if node, ok = d.nodes[row[0]]; !ok {
			node = NewNode(row[0])
			d.nodes[row[0]] = node
		}

		var other *Node
		if other, ok = d.nodes[row[1]]; !ok {
			other = NewNode(row[1])
			d.nodes[row[1]] = other
		}

		if !node.IsEnd() && !other.IsStart() {
			node.Connected = append(node.Connected, other)
		}
		if !other.IsEnd() && !node.IsStart() {
			other.Connected = append(other.Connected, node)
		}
	}

	return nil
}

func pathContains(path string, id string) bool {
	// paths always start with "start," and can't backtrack to start
	return strings.Contains(path, ","+id)
}

func (d *Today) walk(paths []string, currentPath string, currentNode *Node) []string {
	if currentNode.IsEnd() {
		paths = append(paths, currentPath+"end")
	} else if currentNode.IsBig() || !pathContains(currentPath, currentNode.Id) {
		currentPath = currentPath + currentNode.Id + ","
		for _, node := range currentNode.Connected {
			paths = d.walk(paths, currentPath, node)
		}
	}

	return paths
}

// returns:
//  - If the node can be visited
//  - If this path has duplicate small nodes after the addition of this node
func (d *Today) canVisit(path string, hasDuplicate bool, node *Node) (bool, bool) {
	if node.IsBig() || node.IsEnd() || node.IsStart() || !pathContains(path, node.Id) {
		return true, hasDuplicate
	}

	// node is small, path has the node already

	if hasDuplicate {
		return false, hasDuplicate
	}

	return true, true
}

func (d *Today) walk2(paths []string, currentPath string, hasDupe bool, currentNode *Node) []string {
	if currentNode.IsEnd() {
		paths = append(paths, currentPath+"end")
	} else if canVisit, willHaveDupe := d.canVisit(currentPath, hasDupe, currentNode); canVisit {
		currentPath = currentPath + currentNode.Id + ","
		for _, node := range currentNode.Connected {
			paths = d.walk2(paths, currentPath, willHaveDupe, node)
		}
	}

	return paths
}

func (d *Today) Part1() (string, error) {
	paths := make([]string, 0)
	paths = d.walk(paths, "", d.nodes["start"])

	return strconv.Itoa(len(paths)), nil
}

func (d *Today) Part2() (string, error) {
	paths := make([]string, 0)
	paths = d.walk2(paths, "", false, d.nodes["start"])

	return strconv.Itoa(len(paths)), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
