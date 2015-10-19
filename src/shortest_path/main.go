package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Node struct {
	x, y int
	next *Node
}

type Point struct {
	x, y int
}

func NewNode(x, y int) *Node {
	return &Node{x, y, nil}
}

func (n *Node) Distance(node *Node) float64 {
	return math.Sqrt(math.Pow(float64(n.x)-float64(node.x), 2) + math.Pow(float64(n.y)-float64(node.y), 2))
}

func TotalDistance(node *Node) float64 {
	total := 0.0

	currentNode := node
	nextNode := currentNode.next

	for nextNode != nil {
		total += currentNode.Distance(nextNode)

		currentNode = nextNode
		nextNode = currentNode.next
	}

	return total
}

func Nearest(node *Node, x, y int) {
	nearest := math.MaxFloat64

	newNode := NewNode(x, y)
	currentNode := node

	var nearestNode *Node

	for currentNode != nil {
		if value := currentNode.Distance(newNode); value < nearest {
			nearest = value
			nearestNode = currentNode
		}
		currentNode = currentNode.next
	}

	nearestNode.next, newNode.next = newNode, nearestNode.next

}

func Shortest(node *Node, x, y int) {
	nearest := math.MaxFloat64

	newNode := NewNode(x, y)
	currentNode := node
	nextNode := currentNode.next

	var nearestNode *Node

	if nextNode == nil {
		currentNode.next = newNode
		return
	}

	for nextNode != nil {
		if value := currentNode.Distance(newNode) + newNode.Distance(nextNode); value < nearest {
			nearest = value
			nearestNode = currentNode
		}
		currentNode = nextNode
		nextNode = currentNode.next
	}

	nearestNode.next, newNode.next = newNode, nearestNode.next
}

func main() {
	var values []Point

	newSource := rand.NewSource(1235154)
	newRand := rand.New(newSource)

	values = append(values, Point{newRand.Intn(500), newRand.Intn(500)})

	for i := 0; i < 50; i++ {
		values = append(values, Point{newRand.Intn(500), newRand.Intn(500)})
	}

	rootNode := NewNode(values[0].x, values[0].y)
	rootNodeTwo := NewNode(values[0].x, values[0].y)

	values = values[1:]

	for _, value := range values {
		Nearest(rootNode, value.x, value.y)
		Shortest(rootNodeTwo, value.x, value.y)
	}

	currentNode := rootNode

	for currentNode != nil {
		fmt.Printf("[%d, %d]\n", currentNode.x, currentNode.y)
		currentNode = currentNode.next
	}

	fmt.Println(TotalDistance(rootNode))

	fmt.Println("\n")

	currentNode = rootNodeTwo

	for currentNode != nil {
		fmt.Printf("[%d, %d]\n", currentNode.x, currentNode.y)
		currentNode = currentNode.next
	}

	fmt.Println(TotalDistance(rootNodeTwo))
}
