package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
	
	var message string = "Hello Go World!"
	var a, b, c int = 1, 2, 3
	fmt.Println(message, a, b, c)
	
	// you don't need to specify types (in some cases I guess?). They can be infered.
	
	var message2 = "Hello Go World!"
	var d, e, f = 1, 2, 3
	fmt.Println(message2, d, e, f)
	
	// := operator does the declaration and initialization at the same time
	
	message3 := "Hello Go World!"
	g, h, i := 1, 2, 3
	fmt.Println(message3, g, h, i)
}