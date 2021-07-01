/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main

// imported package(if not used shows error)
import "fmt"

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
        
    
    
}
