package main

import "fmt"

func main() {
	r := rectangle{5,6}
	c := circle{7}
	s := square{4}
	fmt.Println("area : ",r.area(), c.area(), s.area())
	fmt.Println("perimeter : ",r.perimeter(),c.perimeter(),s.perimeter())

	shapes := []shape{r,c,s}

	fmt.Println("Hello Amar")
	for _, data := range shapes {
		//fmt.Println(data.area(), data.perimeter())
		fmt.Println(findArea(data))
	}
	/////////////// code 1 ///////////////////////
	//Push(10)
	//Push(20)
	//Push(30)
	//Push(40)
	//Push(50)
	//Print()
	//Reverse()
	//Print()
	/////////////// code 2 /////////////////////
	//li := &LinkedList{}
	//li.Push(10)
	//li.Push(20)
	//li.Push(30)
	//li.Push(40)
	//li.Push(50)
	//li.Printing()
	//li.Reverse()
	//li.Printing()
	//////////////////////////////////////////////


}
/////////////////////////////////////////////////////////////////////////////
//package main
//
//import "fmt"
//
//var vis [1001]int
//
//
//func dfs(src int) {
//	if vis[src]==1 {
//		return
//	}
//	vis[src] = 1
//	fmt.Println(src)
//	for data :
//}
//func main() {
//	for i:=1 ; i<=10 ; i++ {
//		vis[i] = 0
//	}
//}
