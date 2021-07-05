package main
import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	length int
	head  *node
}

func (l *linkedlist) Push(val int) {
	new_node := node{val,nil}
	if l.head == nil {
		l.head = &new_node
		l.length = l.length + 1
	} else {
		new_node.next = l.head
		l.head = &new_node
		l.length = l.length + 1
	}
}

func (l linkedlist) Printing() {
	if l.head == nil {
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data," ")
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	li := &linkedlist{}
	li.Push(10)
	li.Push(20)
	li.Push(30)
	li.Push(40)
	li.Push(50)
	li.Printing()



}

