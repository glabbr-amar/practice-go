/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main

// imported package(if not used shows error)
import "fmt"
import "math/rand"
import "time"


// if something is declared global but not used, it will not provide error 
var x, y, z = 123, true, "foo"

// init function is called before main and executes sequencially
// we can have multiple init funciton
func init() {
	fmt.Println("hi,", "Amar")
}
func main() {
    // use of if-else
    if n := 25; n%2 == 0 {
		fmt.Println(n, "is an even number.")
	} else {
		fmt.Println(n, "is an odd number.")
	}
	
	// we can keep conditional statement in () 
	if 25<23 {
	    println("yes")
	} else {
	    println("No")
	}
	
	for i := 0; i < 3; i++ {
    	fmt.Println(i)
    }
    
    var i = 0
    for ; i < 2; {
    	fmt.Println(i)
    	i++
    }
    for i < 4 {
    	fmt.Println(i)
    	i++
    }
    for i<10 {
        fmt.Println(i)
        if i>=6 {
            break
        }
        i++
    }
    
    // continue and break is similar to C/C++ 
    
    // infinite loop
    /*for ; true; {
    }
    for true {
    }
    for ; ; {
    }
    for {
    }*/
    
    rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n%9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		// here, this "break" statement is nonsense.
		break
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
	
	// in Go break is always there by default. so we need not use break.
	// if we want to remove by default break then we use fallthrough(but should not be used in final case or default case or in the middle of statement of any case)
	switch n := rand.Intn(100) % 5; n {
    case 0, 1, 2, 3, 4:
    	fmt.Println("n =", n)
    	// The "fallthrough" statement makes the
    	// execution slip into the next branch.
    	fallthrough
    case 5, 6, 7, 8:
    	// A new declared variable also called "n",
    	// it is only visible in the currrent
    	// branch code block.
    	n := 99
    	fmt.Println("n =", n) // 99
    	fallthrough
    default:
    	// This "n" is the switch expression "n".
    	fmt.Println("n =", n)
    }
    
    // goto is similar to C/C++
    i = 0
    Next: // here, a label is declared.
    	fmt.Println(i)
    	i++
    	if i < 3 {
    		goto Next // execution jumps
    	}
    
	
}
func init() {
	fmt.Println("hi,", "Kumar")
}
