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



func main() {
    print("Hello",1) // prints without spaces and no endline
    println("Hello",2) // prints with spaces and endline
    fmt.Println("Hello World") // it will be used in newer version but above may be depreciated
    
    // numbers can be written with spaces for better visualization
    /*println(29 == 2_9) prints true*/ 
    
    
    
}
