package main

import (
	"fmt"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func newTreeNode(data string) *TreeNode {
	newNode := new(TreeNode)
	newNode.Data = data
	newNode.Left = nil
	newNode.Right = nil
	return newNode
}

func NewTree(data string) *Tree {
	rootNode := newTreeNode(data)
	return &Tree{rootNode}
}

func (t *Tree) AddNode(data string) {
	newNode := newTreeNode(data)

	currentNode := t.Root

	for {
		if currentNode.Data > data {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		} else if currentNode.Data < data {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		}
	}
}

func (t *Tree) CreateList() []string {
	treeList := make([]string, 0)

	strChan := walkTree(t.Root)

	for {
		str, ok := <-strChan
		if ok == false {
			break
		}
		treeList = append(treeList, str)
	}

	return treeList
}

func (t *Tree) PrintTree() {
	ch := walkTree(t.Root)

	for {
		str, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(str)
	}
}

func walk(node *TreeNode, ch chan string) {
	if node.Left != nil {
		walk(node.Left, ch)
	}
	ch <- node.Data
	if node.Right != nil {
		walk(node.Right, ch)
	}
}

func walkTree(node *TreeNode) <-chan string {
	ch := make(chan string)

	go func() {
		walk(node, ch)
		close(ch)
	}()

	return ch
}

func main() {
	aTree := NewTree("hello")
	aTree.AddNode("females")
	aTree.AddNode("yo")
	aTree.AddNode("awesome")
	aTree.AddNode("dude")
	aTree.AddNode("wax")
	aTree.AddNode("are")
	aTree.AddNode("ant")

	aTree.PrintTree()
	treeList := aTree.CreateList()
	fmt.Println(treeList)
}
