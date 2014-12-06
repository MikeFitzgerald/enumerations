/*

LESSONS

1. Identifiers are case-sensitive.

2. Identifiers that begin with a capital letter are considered to be public (EXPORTED in Go 
   terminology).

3. Identifers that begin with a lower-case letter are consdiered to be private (UNEXPORTED 
   in Go terminology).

4. The blank identifier, _, servers as a placeholder for where a variable is expected in an 
   assignment, and discards any value it is given.

5. The blank identifier is not considered to be a new variable, so if it is used with the := 
   operator, at least one other (new) variable must be assigned too.

6. You can discard some or all of a function's return values by assigning them to the blank 
   identifier.

	Example:
		count, err = fmt.Println(x)	// get number of byes printed and error.
		count, _ = fmt.Println(x)	// get number of bytes printed; discard error.
		_, err = fmt.Println(x)		// discard number of bytes printed; get error.
		fmt.Println(x)				// ignore return values

7. Go offers built-in barebones enumeration support

8. The iota predefined identifier represents successive untyped integer constants; its value is 
   reset to zero whenever the keyword const occurs, and increments by one for each constant 
   declaration.


*/

package main

import (
   "fmt"
   "strings"
)

func main() {

	flag := Active | Receive
	fmt.Println(flag.String())

	// Constants and variables
	// const limit = 512	// constant; type-compatible with any number
	// const top uint16 = 1421	// constant; type: uint16
	// start := -19	// variable; inferred type: int
	// end := int64(9876543210) // variable; type: int64
	// var i int 	//variable; value 0; type: int
	// var debug = false // variable; inferred type: bool
	// checkResults := true // variable; inferred type: bool
	// stepSize := 1.5 // variable; inferred type: float64
	// acronym := "FOSS"	// variable; inferred type: string

	// Enumerations
	// iota:  A jot; a very small, inconsiderable quantity. 
	const Cyan = 0
	const Magenta = 1
	const Yellow = 2

	/*
	These are also equivalent:
	
	const (
		Cyan = 0
		Magenta = 1
		Yellow = 2
	)

	const (
		Cyan = itoa
		Magenta
		Yellow
	)
	*/
}

// Also possible to use iota with floating-point numbers, simple expressions, and custom types.
// In the following snippet, we created three bit flags of custom type BitFlag and then set 
// variable flag (of type BitFlag) to the bitwise OR of the two of them (so flag has value 3).
type BitFlag int

const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // Implicitly BitFlag = 1 << iota // 1 << 1 == 2
	Receive // Implicitly BigFlag = 1 << iota // 1 << 2 == 4
)

// Build up a (possibly empty) slice of strings for those bit
func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 { // int(flag) is vital to avoid infinite recursion!
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}