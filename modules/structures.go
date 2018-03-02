package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	Parent    *Node        `json:"-"`
	Children  []*Node      `json:"children,omitempty"`
	Interface *interface{} `json:"Interface,omitempty"`
	Label     string       `json:"label,omitempty"`
}

func (n *Node) Print() {

	value := ""
	if n.Interface != nil {
		value = fmt.Sprint(*n.Interface)
	}

	fmt.Printf("node label (%s) Parent node (%v) has (%d) children and Interface is (%#v) interface ToString(%s)\n",
		n.Label, n.Parent,
		len(n.Children),
		n.Interface,
		value)

	for indice, node := range n.Children {
		fmt.Printf("node children (%d) from (%s): \t", indice, n.Label)
		node.Print()
	}
}

func NewNodeWithInterface(parent *Node, label string, i interface{}) *Node {
	newNode := NewNode(parent, label)
	newNode.Interface = &i
	return newNode
}

func NewNode(parent *Node, label string) *Node {

	newNode := &Node{
		Children: make([]*Node, 0),
		Label:    label,
	}
	if parent != nil {
		newNode.Parent = parent
		if parent.Children != nil {
			parent.Children = append(parent.Children, newNode)
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

	setParent(node.Children, node)

	return node
}

func setParent(children []*Node, Parent *Node) {
	for _, n := range children {
		n.Parent = Parent
		setParent(n.Children, n)
	}
}

func (n *Node) ReplaceByThisNode(newnode *Node) *Node {
	// set the new Parent pointer
	newnode.Parent = n.Parent

	// remove the srcnode from Parent's children
	children := newnode.Parent.Children
	newnode.Parent.Children = make([]*Node, 0)
	for _, child := range children {
		if child != n {
			newnode.Parent.Children = append(newnode.Parent.Children, child)
		} else {
			// 	add the newnode in Parent's children
			newnode.Parent.Children = append(newnode.Parent.Children, newnode)
		}
	}

	// add all srcnode children to newnode
	for _, child := range n.Children {
		if child != nil {
			newnode.Children = append(newnode.Children, child)
		}
	}

	n.Parent = nil
	n.Children = make([]*Node, 0)

	return newnode
}

func replaceNode(srcNode, newNode *Node) *Node {

	// set the new Parent pointer
	newNode.Parent = srcNode.Parent

	// remove the srcnode from Parent's children
	children := newNode.Parent.Children
	newNode.Parent.Children = make([]*Node, 0)
	for _, child := range children {
		if child != srcNode {
			newNode.Parent.Children = append(newNode.Parent.Children, child)
		} else {
			// 	add the newnode in Parent's children
			newNode.Parent.Children = append(newNode.Parent.Children, newNode)
		}
	}

	// add all srcnode children to newnode
	for _, child := range srcNode.Children {
		if child != nil {
			newNode.Children = append(newNode.Children, child)
		}
	}

	srcNode.Parent = nil
	srcNode.Children = make([]*Node, 0)

	return newNode
}
