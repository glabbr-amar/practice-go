package main

import (
	"fmt"
)

type Book struct {
	title, author string
	pages         int
}

func fib(n int) int {
    if n<=1 {
        return 1
    }
    return fib(n-1) + fib (n-2)
}
func main() {
	book := Book{"Go 101", "Tapir", 256}
	fmt.Println(book) // {Go 101 Tapir 256}

	book = Book{author: "Tapir", pages: 256, title: "Go 101"}
	fmt.Println(book)

	book = Book{}
    fmt.Println(book)

	book = Book{author: "Tapir"}
    fmt.Println(book)

	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir Liu"
	book2.pages = 300
	fmt.Println(book2.pages) // 300
	
	n:=6
	flag := 0
	for i:=2 ; i*i<=n ; i++ {
	    if n%i==0 {
	        flag = 1
	        break
	    }
	}
    // unlike c/C++ we can't use if flag, it needs to be if flag==some_value
	if flag==1 {
	    fmt.Println("not prime")
	}else{
	    fmt.Println("prime")
	}
	
	var res = fib(3)
	fmt.Println("res is : ",res)
}