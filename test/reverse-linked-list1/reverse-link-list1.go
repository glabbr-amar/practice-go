package main

import (
	"log"
)

type Node struct {
	data int
	next *Node
}

var headNode *Node = nil

func push(value int) {
	newNode := &Node{}
	newNode.data = value
	newNode.next = headNode
	headNode = newNode
}

func print() {
	defer log.Println("printing done")
	if headNode != nil {
		var iterator *Node = headNode
		for iterator != nil {
			log.Print(iterator.data, " ")
			iterator = iterator.next
		}
	}
}

func reverse() {
	var previousNode *Node = nil
	var nextNode *Node = nil
	var currentNode *Node = headNode
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	headNode = previousNode
}
