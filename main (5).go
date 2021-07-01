/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main

// imported package(if not used shows error)
import "fmt"

// multiple package can be imported like this
import (
	"math/rand"
	"time"
)
// import syntax
// if import name is given as some name, it will be used by that name
// import importname "path/to/package"
// if import name is . then we need can use the function directly without specifying fucntion
// example-> import . "fmat"
// we can use Println() directly

import _ "math/rand" // okay: it is not required to be used

// use of const
// here Yes is used before its declaration but this will give error when used inside main function
const n = 3.1416
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
)

// if something is declared global but not used, it will not provide error 
var x, y, z = 123, true, "foo"

// init function is called before main and executes sequencially
// we can have multiple init funciton
func init() {
	fmt.Println("hi,", "Amar")
}
func main() {
    print("Hello",1) // prints without spaces and no endline
    println("Hello",2) // prints with spaces and endline
    fmt.Println("Hello World") // it will be used in newer version but above may be depreciated
    
    // numbers can be written with spaces for better visualization
    /*println(29 == 2_9) prints true*/ 
    
    // Y and Z will take the value of its previous initialized variable, similary C and  _ will take the value of A and B but _ can't be used to print 
    const (
    	X float32 = 3.14
    	Y           
    	Z           
    
    	A, B = "Go", "language"
    	C, _
    	
    )
    
    // variable if declared but not used will provide error
    /*var lang, website string = "Go", "https://golang.org"
    var compiled, dynamic bool = true, false
    var announceYear int = 2009
    var ( 
        lang, bornYear, compiled     = "Go", 2007, true
    	announceAt, releaseAt    int = 2009, 2012
    	createdBy, website       string
    )*/
    
    // short declaration of variables(also provide error if not used)
    /*lang, year := "Go language", 2007*/
        
    // binary operators are similar to C/C++, but in unary operator we use ^ instead of ~ for negation
    
    // anonymous function
    x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	func(a, b int) {
		println("a*a + b*b =", a*a + b*b)
	}(x, y) // pass argument x and y to parameter a and b.
    
    
    // use of random package
    rand.Seed(time.Now().UnixNano())
	var temp1 = rand.Int31()
	var temp2 = rand.Float32()
	var temp3 = rand.Int31n(20); // takes integer below 20
	println(temp1,temp2,temp3)
    
}
func init() {
	fmt.Println("hi,", "Kumar")
}
// use of function
func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a + b, a - b
	s = x * x
	d = y * y
	return // <=> return s, d
}

func SquaresOfSumAndDiffs(a int64, b int64) (int64, int64) {
	return (a+b) * (a+b), (a-b) * (a-b)
}
