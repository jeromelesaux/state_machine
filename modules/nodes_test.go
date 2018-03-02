package modules

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCreationparentNode(t *testing.T) {
	parent := NewNode(nil, "parent node 0")
	fmt.Println(parent)
}

func TestCreateChildren(t *testing.T) {
	parent := NewNode(nil, "parent node 0")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	childrenNumber := r.Intn(100)
	fmt.Printf("Number %d\n", childrenNumber)
	for i := 0; i <= childrenNumber; i++ {
		newNode := NewNode(parent, "Node"+strconv.Itoa(i))
		fmt.Println(newNode)
	}

	subNode := NewNode(parent.Children[0], "Node00")
	fmt.Println(subNode)
	fmt.Println(parent)
}

func TestSerializeGraph(t *testing.T) {
	parent := NewNode(nil, "parent node 0")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	childrenNumber := r.Intn(100)
	for i := 0; i <= childrenNumber; i++ {
		NewNode(parent, "Node"+strconv.Itoa(i))
	}
	NewNode(parent.Children[0], "Node00")
	fmt.Println(parent)
	f, _ := os.Create("graph_test.json")
	err := json.NewEncoder(f).Encode(parent)
	if err != nil {
		fmt.Println("Error " + err.Error())
		t.FailNow()
	}
}

func TestDeserializeGrah(t *testing.T) {
	node := NewNode(nil, "")
	f, _ := os.Open("graph_test.json")
	err := json.NewDecoder(f).Decode(node)
	if err != nil {
		fmt.Println("Error " + err.Error())
		t.FailNow()
	}

	fmt.Println(node.Children[0])
}

func TestLoadNodeFromFilepath(t *testing.T) {
	node := LoadNodes("graph_test.json")
	if node.Children[0].Parent == nil {
		t.FailNow()
	}
	fmt.Println(node.Children[0])
}

func TestReplaceNode(t *testing.T) {
	parentNode := NewNode(nil, "root node")
	srcNode := NewNode(parentNode, "source node")
	newNode := NewNode(nil, "new node")
	parentNode.Print()
	replaceNode(srcNode, newNode)
	parentNode.Print()
	if parentNode.Children[0] != newNode {
		t.Fatalf("expetected node %v and results %v", newNode, parentNode.Children[0])
	}
}

func TestReplaceFunctionNode(t *testing.T) {
	parentNode := NewNode(nil, "root node")
	srcNode := NewNode(parentNode, "source node")
	newNode := NewNode(nil, "new node")
	parentNode.Print()
	srcNode.ReplaceByThisNode(newNode)
	parentNode.Print()
	if parentNode.Children[0] != newNode {
		t.Fatalf("expetected node %v and results %v", newNode, parentNode.Children[0])
	}
}

type message struct {
	M string
}

func (m *message) ToString() string {
	return m.M
}

func TestInterfaceUsage(t *testing.T) {

	m1 := &message{M: "hello world message 1"}
	parentNode := NewNodeWithInterface(nil, "root node", m1)
	m2 := &message{M: "hello world message 2"}
	NewNodeWithInterface(parentNode, "source node", m2)
	m3 := &message{M: "hello world message 3"}
	NewNodeWithInterface(parentNode, "new node", m3)
	if len(parentNode.Children) != 2 {
		t.Fatalf("expected length 2 and gets %d", len(parentNode.Children))
	}
	parentNode.Print()
}
