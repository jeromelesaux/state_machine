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

func TestCreationMotherNode(t *testing.T) {
	mother := NewNode(nil, "mother node 0")
	fmt.Println(mother)
}

func TestCreateChildren(t *testing.T) {
	mother := NewNode(nil, "mother node 0")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	childrenNumber := r.Intn(100)
	fmt.Printf("Number %d\n", childrenNumber)
	for i := 0; i <= childrenNumber; i++ {
		newNode := NewNode(mother, "Node"+strconv.Itoa(i))
		fmt.Println(newNode)
	}

	subNode := NewNode(mother.Children[0], "Node00")
	fmt.Println(subNode)
	fmt.Println(mother)
}

func TestSerializeGraph(t *testing.T) {
	mother := NewNode(nil, "mother node 0")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	childrenNumber := r.Intn(100)
	for i := 0; i <= childrenNumber; i++ {
		NewNode(mother, "Node"+strconv.Itoa(i))
	}
	NewNode(mother.Children[0], "Node00")
	fmt.Println(mother)
	f, _ := os.Create("graph_test.json")
	err := json.NewEncoder(f).Encode(mother)
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
	if node.Children[0].Mother == nil {
		t.FailNow()
	}
	fmt.Println(node.Children[0])
}

func TestReplaceNode(t *testing.T) {
	motherNode := NewNode(nil, "root node")
	srcNode := NewNode(motherNode, "source node")
	newNode := NewNode(nil, "new node")
	motherNode.Print()
	replaceNode(srcNode, newNode)
	motherNode.Print()
	if motherNode.Children[0] != newNode {
		t.Fatalf("expetected node %v and results %v", newNode, motherNode.Children[0])
	}
}

func TestReplaceFunctionNode(t *testing.T) {
	motherNode := NewNode(nil, "root node")
	srcNode := NewNode(motherNode, "source node")
	newNode := NewNode(nil, "new node")
	motherNode.Print()
	srcNode.ReplaceByThisNode(newNode)
	motherNode.Print()
	if motherNode.Children[0] != newNode {
		t.Fatalf("expetected node %v and results %v", newNode, motherNode.Children[0])
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
	motherNode := NewNodeWithInterface(nil, "root node", m1)
	m2 := &message{M: "hello world message 2"}
	NewNodeWithInterface(motherNode, "source node", m2)
	m3 := &message{M: "hello world message 3"}
	NewNodeWithInterface(motherNode, "new node", m3)
	if len(motherNode.Children) != 2 {
		t.Fatalf("expected length 2 and gets %d", len(motherNode.Children))
	}
	motherNode.Print()
}
