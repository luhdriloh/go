package main

import "fmt"

type Node struct {
	Former *Node
	Next   *Node
	Data   int
}

func CreateNode(data int) *Node {
	return &Node{nil, nil, data}
}

func (n *Node) AddNode(data int) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &Node{n, nil, data}
}

func (n *Node) DeleteNode(node_to_delete *Node) {
	for n != nil {
		if n == node_to_delete {
			n.Former.Next = n.Next
			break
		}
		n = n.Next
	}
}

func (n *Node) ReturnData() int {
	return n.Data
}

func main() {
	root_node := CreateNode(10)
	for i := 0; i < 10; i++ {
		root_node.AddNode(i * 7)
	}

	current_node := root_node

	for current_node != nil {
		fmt.Println(current_node.ReturnData())
		current_node = current_node.Next
	}

	fmt.Println("\n")

	node_to_delete := root_node.Next.Next
	root_node.DeleteNode(node_to_delete)
	current_node = root_node

	for current_node != nil {
		fmt.Println(current_node.ReturnData())
		current_node = current_node.Next
	}
}
