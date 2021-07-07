package main

func main() {
	linkList := &LinkedList{}
	linkList.push(10)
	linkList.push(20)
	linkList.push(30)
	linkList.push(40)
	linkList.push(50)
	linkList.print()
	linkList.reverse()
	linkList.print()
}
