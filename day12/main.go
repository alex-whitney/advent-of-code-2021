package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2021/lib"
)

type Node struct {
	Id        string
	Connected []*Node
}

func NewNode(id string) *Node {
	return &Node{
		Id:        id,
		Connected: []*Node{},
	}
}

func (node *Node) IsBig() bool {
	return strings.ToUpper(node.Id) == node.Id
}

func (node *Node) IsEnd() bool {
	return node.Id == "end"
}

func (node *Node) IsStart() bool {
	return node.Id == "start"
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

func (d *Today) canVisit(path string, node *Node) bool {
	if node.IsBig() || node.IsEnd() || node.IsStart() || !pathContains(path, node.Id) {
		return true
	}

	// look for two small caves already on path - otherwise, this small node can be added
	counts := make(map[string]int)
	ids := strings.Split(path, ",")
	for _, id := range ids[:len(ids)-1] {
		// start and end aren't big, but appear at most once
		if !d.nodes[id].IsBig() {
			counts[id]++
			if counts[id] > 1 {
				return false
			}
		}
	}
	return true
}

func (d *Today) walk2(paths []string, currentPath string, currentNode *Node) []string {
	if currentNode.IsEnd() {
		paths = append(paths, currentPath+"end")
	} else if d.canVisit(currentPath, currentNode) {
		currentPath = currentPath + currentNode.Id + ","
		for _, node := range currentNode.Connected {
			paths = d.walk2(paths, currentPath, node)
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
	paths = d.walk2(paths, "", d.nodes["start"])

	return strconv.Itoa(len(paths)), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
