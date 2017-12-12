package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	Mother    *Node        `json:"-"`
	Children  []*Node      `json:"children,omitempty"`
	Interface *interface{} `json:"Interface,omitempty"`
	Label     string       `json:"label,omitempty"`
}

func (n *Node) Print() {

	value := ""
	if n.Interface != nil {
		value = fmt.Sprint(*n.Interface)
	}

	fmt.Printf("node label (%s) mother node (%v) has (%d) children and Interface is (%#v) interface ToString(%s)\n",
		n.Label, n.Mother,
		len(n.Children),
		n.Interface,
		value)

	for indice, node := range n.Children {
		fmt.Printf("node children (%d) from (%s): \t", indice, n.Label)
		node.Print()
	}
}

func NewNodeWithInterface(mother *Node, label string, i interface{}) *Node {
	newNode := NewNode(mother, label)
	newNode.Interface = &i
	return newNode
}

func NewNode(mother *Node, label string) *Node {

	newNode := &Node{
		Children: make([]*Node, 0),
		Label:    label,
	}
	if mother != nil {
		newNode.Mother = mother
		if mother.Children != nil {
			mother.Children = append(mother.Children, newNode)
		}
	}

	return newNode
}

func (n *Node) IsLastChild() bool {

	if n.Children == nil {
		return true
	}
	if len(n.Children) == 0 {
		return true
	}
	return false
}

func LoadNodes(filepath string) *Node {
	node := NewNode(nil, "")
	f, _ := os.Open(filepath)
	err := json.NewDecoder(f).Decode(node)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return node
	}

	setMother(node.Children, node)

	return node
}

func setMother(children []*Node, mother *Node) {
	for _, n := range children {
		n.Mother = mother
		setMother(n.Children, n)
	}
}

func replaceNode(srcNode, newNode *Node) *Node {

	// set the new mother pointer
	newNode.Mother = srcNode.Mother

	// remove the srcnode from mother's children
	children := newNode.Mother.Children
	newNode.Mother.Children = make([]*Node, 0)
	for _, child := range children {
		if child != srcNode {
			newNode.Mother.Children = append(newNode.Mother.Children, child)
		} else {
			// 	add the newnode in mother's children
			newNode.Mother.Children = append(newNode.Mother.Children, newNode)
		}
	}

	// add all srcnode children to newnode
	for _, child := range srcNode.Children {
		if child != nil {
			newNode.Children = append(newNode.Children, child)
		}
	}

	srcNode.Mother = nil
	srcNode.Children = make([]*Node, 0)

	return newNode
}
