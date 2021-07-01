package main

import (
	"fmt"
)

type Node struct{
    data int
    next *Node
}

func push(val int) {
    
    
    
}
func main() {
	var a [5]int
    fmt.Println( a)
    // unlike c/c++, we can't use ++i in for loop
    for i:=0 ; i<5 ; i++ {
        a[i] = i
    }
    fmt.Println(a)
    
    // Initializing an array
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Initializing b : ",b)
    
    // array of string
    t := []string{"g", "h", "ijk"}
    fmt.Println(t)
    
    // creates an empty string s of size 3
    s := make([]string, 3)
    fmt.Println(s)
    
    // appending some letters
    s = append(s,"a","m","ar")
    fmt.Println(s)
    // we can also slice a string similar to python
    // we can copy one string to another using copy function
    // syntax is copy(destination, source)
    copied_string := make([]string,len(s))
    copy(copied_string,s)
    fmt.Println("Copied string is : ",copied_string)
}