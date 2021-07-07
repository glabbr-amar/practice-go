package main

import (
	"log"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	length   int
	headNode *Node
}

func (l *LinkedList) reverse() {
	var previousNode *Node = nil
	var nextNode *Node = nil
	currentNode := l.headNode
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	l.headNode = previousNode
}
func (l *LinkedList) push(val int) {
	newNode := Node{val, nil}
	if l.headNode == nil {
		l.headNode = &newNode
		l.length = l.length + 1
	} else {
		newNode.next = l.headNode
		l.headNode = &newNode
		l.length = l.length + 1
	}
}

func (l *LinkedList) print() {
	if l.headNode == nil {
		return
	}
	currentNode := l.headNode
	for currentNode != nil {
		log.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	log.Println("done printing")
}
