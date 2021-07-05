package main

import "fmt"

type Node struct {
	data int
	next *Node
}

//////////////// code 1 /////////////////////////////
var head *Node = nil
func Push(ele int) {
	temp := &Node{}
	temp.data = ele
	temp.next = head
	head = temp
}

func Print() {
	if head != nil {
		var temp *Node = head
		for temp!=nil {
			fmt.Print(temp.data," ")
			temp = temp.next
		}
		fmt.Println()
	}
}

func Reverse() {
	var prev *Node = nil
	var nxt *Node = nil
	var curr *Node = head
	for curr!=nil {
		nxt = curr.next
		curr.next = prev
		prev = curr
		curr = nxt
	}
	head = prev
}
////////////////// code 2 /////////////////////////////////

//type LinkedList struct {
//	length int
//	head  *Node
//}
//
//func (l *LinkedList) Reverse(){
//	//prev := &Node
//	var prev *Node = nil
//	var nxt *Node = nil
//	curr := l.head
//	for curr!= nil {
//		nxt = curr.next
//		curr.next = prev
//		prev = curr
//		curr = nxt
//	}
//	l.head = prev
//}
//func (l *LinkedList) Push(val int) {
//	newNode := Node{val,nil}
//	if l.head == nil {
//		l.head = &newNode
//		l.length = l.length + 1
//	} else {
//		newNode.next = l.head
//		l.head = &newNode
//		l.length = l.length + 1
//	}
//}
//
//func (l *LinkedList) Printing() {
//	if l.head == nil {
//		return
//	}
//	curr := l.head
//	for curr != nil {
//		fmt.Print(curr.data," ")
//		curr = curr.next
//	}
//	fmt.Println("done printing")
//}


